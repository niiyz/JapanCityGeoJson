FROM osgeo/gdal:alpine-small-latest as make_data

RUN apk add --update --no-cache alpine-sdk tree

RUN mkdir -p /app

WORKDIR /app

RUN gdalinfo --version

# Download data
ADD https://nlftp.mlit.go.jp/ksj/gml/data/N03/N03-2020/N03-20200101_GML.zip .

# Unzip Download File

RUN unzip N03-20200101_GML.zip

RUN tree .

RUN ls N03-200101_GML/N03-20_200101.shp
RUN ogrinfo N03-200101_GML/N03-20_200101.shp

# Shapefile shit_jis -> UTF-8 , Input dataset open option shift_jis, Layer creation option UTF-8
RUN ogr2ogr -f "ESRI Shapefile" -lco ENCODING=UTF-8 -oo ENCODING=shift_jis N03-20_200101sjis.shp /app/N03-200101_GML/N03-20_200101.shp

# Shapefile -> GeoJson
RUN ogr2ogr -f GeoJSON -nlt MULTIPOLYGON JAPAN2020-MULTIPOLYGON.json N03-20_200101sjis.shp

FROM osgeo/gdal:alpine-small-latest as dev

RUN apk add --update --no-cache alpine-sdk git go

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

COPY . /go/src/app

COPY --from=make_data /app/JAPAN2020-MULTIPOLYGON.json .

RUN go version

RUN ogr2ogr

