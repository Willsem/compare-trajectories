# File Parser

Парсер данных, считанных с блоков-датчиков в картинге

## Использование

```
python3 main.py --acc AcceletometerFile --gps GpsFile --out OutFile
```

### Accelerometer

```
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

```
{
    "date": datetime[],
    "long": float[],
    "lat": float[]
}
```
