from datetime import datetime


class Point:
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z


class AccelerometerData:
    def __init__(self):
        self.date = []
        self.acc = []
        self.gyro = []

    def append(self, strdate, ax, ay, az, gx, gy, gz):
        self.date.append(strdate)
        self.acc.append(Point(ax, ay, az).__dict__)
        self.gyro.append(Point(gx, gy, gz).__dict__)


class GpsData:
    def __init__(self):
        self.date = []
        self.lat = []
        self.long = []

    def append(self, strdate, long, lat):
        self.date.append(strdate)

        ld = long // 100
        lm = long % 100
        self.long.append(ld + lm / 60)

        ld = lat // 100
        lm = lat % 100
        self.lat.append(ld + lm / 60)


def get_dict(gps, acc):
    return {
        'gps': gps,
        'acc': acc,
    }


def parse_acc(file):
    parsed_data = AccelerometerData()

    for line in file:
        line = line[:-1] \
            .replace('acc: ', '') \
            .replace('gyro: ', '') \
            .replace('[', '') \
            .replace(']', ',') \
            .replace('; ', ', ') \
            .split(', ')
        parsed_data.append(*line)

    return parsed_data.__dict__


def parse_gps(file):
    parsed_data = GpsData()

    for line in file:
        line = line[:-1] \
            .replace('[', '') \
            .replace('] ', ',') \
            .split(',')

        if line[1] == '$GNRMC':
            try:
                args = [line[0], float(line[4]), float(line[6])]
                parsed_data.append(*args)
            except ValueError:
                pass

    return parsed_data.__dict__


def cut(gps, acc, start, finish):
    if start is None and finish is None:
        return gps, acc
    if start is None:
        start = 0
    if finish is None:
        finish = len(gps['date'])
    gps = {
        'date': gps['date'][start:finish],
        'long': gps['long'][start:finish],
        'lat':  gps['lat'][start:finish],
    }

    def convert_date(string):
        return datetime.strptime(string, '%m/%d/%y %H:%M:%S.%f')

    begin = convert_date(gps['date'][0])
    end = convert_date(gps['date'][-1])

    min_beg = -1
    min_end = -1
    min_beg_i = -1
    min_end_i = -1

    for i in range(len(acc['date'])):
        dist_b = abs(convert_date(acc['date'][i]) - begin).microseconds
        dist_e = abs(convert_date(acc['date'][i]) - end).microseconds

        if min_beg == -1 or dist_b < min_beg:
            min_beg = dist_b
            min_beg_i = i

        if min_end == -1 or dist_e < min_end:
            min_end = dist_e
            min_end_i = i

    acc = {
        'date': acc['date'][min_beg_i:min_end_i],
        'gyro': acc['gyro'][min_beg_i:min_end_i],
        'acc': acc['acc'][min_beg_i:min_end_i],
    }

    return gps, acc


def parse(gpsfile, accfile, start, finish):
    gps = parse_gps(gpsfile)
    acc = parse_acc(accfile)
    gps, acc = cut(gps, acc, start, finish)
    return get_dict(gps, acc)
