# File Parser

Парсер данных, считанных с блоков-датчиков в картинге

## Использование

```sh
python3 main.py --acc AcceletometerFile --gps GpsFile --out OutFile
```

### Accelerometer

```json
{
    "date": datetime[],
    "acc" : [{
        "x": int,
        "y": int,
        "z": int
    }],
    "gyro": [{
        "x": int,
        "y": int,
        "z": int
    }]
}
```

### GPS

```json
{
    "date": datetime[],
    "long": float[],
    "lat": float[]
}
```
