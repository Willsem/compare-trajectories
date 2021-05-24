import sys
import argparse
import json
from parser import parse


def create_argument_parser():
    parser = argparse.ArgumentParser()
    parser.add_argument('-a', '--acc', required=True)
    parser.add_argument('-g', '--gps', required=True)
    parser.add_argument('-o', '--out', required=True)
    return parser


def parse_files(gps_filename, acc_filename, out):
    with open(gps_filename) as gpsfile:
        with open(acc_filename) as accfile:
            data = parse(gpsfile, accfile)

    with open(out, 'w') as file:
        file.write(json.dumps(data))


def main():
    parser = create_argument_parser()
    args = parser.parse_args(sys.argv[1:])

    parse_files(args.gps, args.acc, args.out)
    print(f'Create file {args.out}')


if __name__ == '__main__':
    main()
