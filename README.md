
```bash
curl -LO http://downloads.dbpedia.org/2016-10/core-i18n/ja/geo_coordinates_ja.tql.bz2
bzip2 -d geo_coordinates_ja.tql.bz2
```

```bash
cd main_updator && go build -o main
cat ../test_datas/geo_coordinates_ja.tql | INDEX_LEVEL=15 GOOGLE_APPLICATION_CREDENTIALS=../wm-minilla-89a6be781c74.json ./main

cd main_web && go build -o main
GOOGLE_APPLICATION_CREDENTIALS=../wm-minilla-920a0d698939.json ALLOWED_ORIGINS="http://localhost:4200" ./main
```