
```bash
curl -LO http://downloads.dbpedia.org/2016-10/core-i18n/ja/geo_coordinates_ja.tql.bz2
bzip2 -d geo_coordinates_ja.tql.bz2
```

```bash
cd main_updator && go build -o main
./main
```