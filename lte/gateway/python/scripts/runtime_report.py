#!/usr/bin/env python3

"""
Copyright 2022 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

import argparse
import glob
import logging
import os
import xml.etree.ElementTree as ET
from datetime import datetime
from pathlib import Path
from pprint import pprint


def merge_all_report(list_xml_report_paths, output_path):
    """ 
    Given the paths of all xml report files, merge them all into a single report files
    Args:
        list_xml_report_paths: list of all the report xml files
        output_path : path of output report file in xml format 
    """
    testsuites_node = ET.Element("testsuites")

    testsuites_node.set("name", "AllTests")
    num_all_failures = 0
    num_all_disabled = 0
    num_all_errors = 0
    total_time = 0
    num_all_tests = 0
    init_time = datetime.max
    c = 0

    # iterate through files
    for xml_file_path in list_xml_report_paths:
        test_result = ET.parse(xml_file_path)
        test_suites_data = test_result.getroot()

        num_all_failures += int(test_suites_data.attrib['failures'])
        num_all_tests += int(test_suites_data.attrib['tests'])
        total_time += float(test_suites_data.attrib['time'])
        num_all_errors += int(test_suites_data.attrib['errors'])
        num_all_disabled += int(test_suites_data.attrib['disabled'])
        init_time = min(init_time, datetime.fromisoformat(test_suites_data.attrib['timestamp']))

        for single_test_suite_data in test_suites_data:
            testsuites_node.append(single_test_suite_data)

    # adding node here
    testsuites_node.set("timestamp", str(init_time))
    testsuites_node.set("time", str(total_time))
    testsuites_node.set("errors", str(num_all_errors))
    testsuites_node.set("disabled", str(num_all_disabled))
    testsuites_node.set("failures", str(num_all_failures))
    testsuites_node.set("tests", str(num_all_tests))

    tree = ET.ElementTree()
    tree._setroot(testsuites_node)
    tree.write(output_path)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)

    magma_root_path = os.environ.get('MAGMA_ROOT')
    default_output_path = ((magma_root_path + "/report/") if magma_root_path else "/tmp/magma_oai/report/") + "report_all_tests.xml"

    parser = argparse.ArgumentParser(description="Merging all the .xml unittest reports into a single file")
    parser.add_argument('-i', '--input', nargs='+', help='<Required> Path of xml files', required=True)
    parser.add_argument('-o', '--output', help='Output of xml report', default=default_output_path, required=False)
    args = parser.parse_args()

    print("=" * 50)

    if len(args.input) < 1:
        logging.info("No report is generated")
    else:
        Path(args.output).parent.mkdir(parents=True, exist_ok=True)
        merge_all_report(args.input, args.output)
    logging.info(f"Final report is generated at: {Path(args.output).resolve()}")
