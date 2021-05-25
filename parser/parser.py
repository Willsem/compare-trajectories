class Point:
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z


class AccelerometerData:
    def __init__(self):
        self.date = []
        self.accelerometer = []
        self.gyroscope = []

    def append(self, strdate, ax, ay, az, gx, gy, gz):
        self.date.append(strdate)
        self.accelerometer.append(Point(ax, ay, az).__dict__)
        self.gyroscope.append(Point(gx, gy, gz).__dict__)


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


def parse(gpsfile, accfile):
    return get_dict(
        parse_gps(gpsfile),
        parse_acc(accfile),
    )
