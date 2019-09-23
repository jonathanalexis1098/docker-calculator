## Getting Started

1. Install dependencies.
```pip
pip3 install -r requirements.txt
```
2. Run app.py
```bash
python3 src/app.py
```

## Dockerizing

1. Create image
```bash
docker build -t "imageName" .
```

2. Run container
```bash
docker run -it -p 3000:4000 "imageName"
```
