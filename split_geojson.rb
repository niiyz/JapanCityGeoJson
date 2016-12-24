require 'json'
require 'fileutils'

class GeoJsonToCity

  def initialize()
    @datas = {}
  end

  def split(path)
    Dir.glob(path).each do |path|
      str      = File.open(path).read(nil, '')
      json     = JSON.parse(str)
      features = json['features']
      features.each do |feature|
        # 名称抽出
        addr = self.parseProperties(feature['properties'])
        # properties削除
        feature.delete('properties')
        # 都道府県
        self.setData('47都道府県', addr[0], feature)
        # 市町村1 ex. ["東京都", "府中市"] , ["広島県", "府中市"]
        self.setData(addr[0], addr[1], feature)
        # 市町村2 ex.["沖縄県", "八重山郡", "与那国町"]
        if !addr[2].nil?
          self.setData(addr[0], addr[1] + addr[2], feature)
        end
      end
    end
  end

  def setData(dirName, fileName, feature)
    key = dirName + '-' + fileName
    if !key.nil?
      if !@datas.has_key?(key) then
        @datas[key] = []
      end
      feature['id'] = fileName
      @datas[key].push(feature)
    end
  end

  def parseProperties(prop)
    addr = []
    pref  = prop['N03_001']
    city1 = prop['N03_002']
    city2 = prop['N03_003']
    city3 = prop['N03_004']
    addr.push(pref)
    if !city1.nil? && pref != '北海道'
      addr.push(city1)
    end
    if !city2.nil?
      addr.push(city2)
    end
    if !city3.nil?
      addr.push(city3)
    end
    return addr
  end

  def make
    @datas.each do |key, collection|
      data = {'type' => 'FeatureCollection', 'features' => []}
      data[:features] = collection
      p key
      prefDir, fileName = key.split('-')
      FileUtils.mkdir_p("geojson/#{prefDir}")
      File.open("geojson/#{prefDir}/#{fileName}.json", 'w').write(JSON.generate(data))
    end
  end

end

city = GeoJsonToCity.new
japanGeoJsonAll = './data/geojson/japan2016.json'
city.split(japanGeoJsonAll)
city.make()
