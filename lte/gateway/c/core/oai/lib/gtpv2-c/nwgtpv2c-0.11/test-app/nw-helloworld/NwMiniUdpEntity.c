/*----------------------------------------------------------------------------*
 *                                                                            *
              M I N I M A L I S T I C     U D P     E N T I T Y
 *                                                                            *
                      Copyright (C) 2010 Amit Chawre.
 *                                                                            *
  ----------------------------------------------------------------------------*/

/**
   @file NwMiniUdpEntity.c
   @brief This file contains example of a minimalistic ULP entity.
*/

#include <stdio.h>
#include <string.h>
#include <assert.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <errno.h>
#include <fcntl.h>
#include "NwEvt.h"
#include "NwGtpv2c.h"
#include "NwMiniLogMgrEntity.h"
#include "NwMiniUdpEntity.h"

#ifndef NW_ASSERT
#define NW_ASSERT assert
#endif

#define MAX_UDP_PAYLOAD_LEN (4096)

#ifdef __cplusplus
extern "C" {
#endif

static NwCharT* gLogLevelStr[] = {"EMER", "ALER", "CRIT", "ERRO",
                                  "WARN", "NOTI", "INFO", "DEBG"};

/*---------------------------------------------------------------------------
   Private functions
  --------------------------------------------------------------------------*/

static void NW_EVT_CALLBACK(nwUdpDataIndicationCallbackData) {
  nw_rc_t rc;
  uint8_t udpBuf[MAX_UDP_PAYLOAD_LEN];
  NwS32T bytesRead;
  uint32_t peerLen;
  struct sockaddr_in peer;
  NwGtpv2cNodeUdpT* thiz = (NwGtpv2cNodeUdpT*)arg;

  peerLen = sizeof(peer);
  bytesRead = recvfrom(fd, udpBuf, MAX_UDP_PAYLOAD_LEN, 0,
                       (struct sockaddr*)&peer, (socklen_t*)&peerLen);

  if (bytesRead) {
    uint32_t peerIp = (peer.sin_addr.s_addr);

    NW_LOG(NW_LOG_LEVEL_DEBG,
           "Received UDP message of size %u from peer %u.%u.%u.%u:%u",
           bytesRead, (peerIp & 0x000000ff), (peerIp & 0x0000ff00) >> 8,
           (peerIp & 0x00ff0000) >> 16, (peerIp & 0xff000000) >> 24,
           ntohs(peer.sin_port));
    rc = nwGtpv2cProcessUdpReq(thiz->hGtpv2cStack, udpBuf, bytesRead,
                               ntohs(peer.sin_port),
                               ntohl(peer.sin_addr.s_addr));
  } else {
    NW_LOG(NW_LOG_LEVEL_ERRO, "%s", strerror(errno));
  }
}

/*---------------------------------------------------------------------------
   Public functions
  --------------------------------------------------------------------------*/

nw_rc_t nwGtpv2cUdpInit(NwGtpv2cNodeUdpT* thiz,
                        nw_gtpv2c_StackHandleT hGtpv2cStack,
                        uint8_t* ipv4Addr) {
  int sd;
  struct sockaddr_in addr;

  sd = socket(AF_INET, SOCK_DGRAM, 0);

  if (sd < 0) {
    NW_LOG(NW_LOG_LEVEL_ERRO, "%s", strerror(errno));
    NW_ASSERT(0);
  }

  addr.sin_family = AF_INET;
  addr.sin_port = htons(2123);
  addr.sin_addr.s_addr = inet_addr(ipv4Addr);
  memset(addr.sin_zero, '\0', sizeof(addr.sin_zero));

  if (bind(sd, (struct sockaddr*)&addr, sizeof(addr)) < 0) {
    NW_LOG(NW_LOG_LEVEL_ERRO, "%s", strerror(errno));
    NW_ASSERT(0);
  }

  NW_EVENT_ADD((thiz->ev), sd, nwUdpDataIndicationCallbackData, thiz,
               NW_EVT_READ | NW_EVT_PERSIST);
  thiz->hSocket = sd;
  thiz->hGtpv2cStack = hGtpv2cStack;
  return NW_OK;
}

nw_rc_t nwGtpv2cUdpDestroy(NwGtpv2cNodeUdpT* thiz) { close(thiz->hSocket); }

nw_rc_t nwGtpv2cUdpDataReq(nw_gtpv2c_UdpHandleT udpHandle, uint8_t* dataBuf,
                           uint32_t dataSize, uint32_t peerIp,
                           uint32_t peerPort) {
  struct sockaddr_in peerAddr;
  NwS32T bytesSent;
  NwGtpv2cNodeUdpT* thiz = (NwGtpv2cNodeUdpT*)udpHandle;

  NW_LOG(NW_LOG_LEVEL_DEBG,
         "Sending buf of size %u for on handle %x to peer %u.%u.%u.%u:%u",
         dataSize, udpHandle, (peerIp & 0xff000000) >> 24,
         (peerIp & 0x00ff0000) >> 16, (peerIp & 0x0000ff00) >> 8,
         (peerIp & 0x000000ff), peerPort);
  peerAddr.sin_family = AF_INET;
  peerAddr.sin_port = htons(peerPort);
  peerAddr.sin_addr.s_addr = htonl(peerIp);
  memset(peerAddr.sin_zero, '\0', sizeof(peerAddr.sin_zero));
  bytesSent = sendto(thiz->hSocket, dataBuf, dataSize, 0,
                     (struct sockaddr*)&peerAddr, sizeof(peerAddr));

  if (bytesSent < 0) {
    NW_LOG(NW_LOG_LEVEL_ERRO, "%s", strerror(errno));
    NW_ASSERT(0);
  }

  return NW_OK;
}

#ifdef __cplusplus
}
#endif
