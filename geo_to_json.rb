require 'json'
require 'fileutils'

class GeoJsonToCity

  def initialize()
   @datas = {}
   Dir.glob('places.json').each do |path|
      io = File.open(path)
      str = io.read(nil, '')
      json = JSON.parse(str)
      features = json['features']
      features.each do |feature|
        # 名称抽出
        city = self.checkCity(feature['properties'])
        self.setData(city[0], feature)
        # 市町村1
        self.setData(city[1], feature)
        # 市町村2
        self.setData(city[2], feature)
        # 市町村3
        self.setData(city[3], feature)
      end
    end
  end

  def setData(key, feature)
    if !key.nil?
      if !@datas.has_key?(key) then
        @datas[key] = []
      end
      @datas[key].push(feature)
    end
  end

  def checkCity(prop)
    cityList = []
    pref  = prop['N03_001']
    city1 = prop['N03_002']
    city2 = prop['N03_003']
    city3 = prop['N03_004']
    cityList.push(pref)
    if !city1.nil? && pref != '北海道'
      cityList.push(city1)
    end
    if !city2.nil?
      cityList.push(city2)
    end
    if !city3.nil?
      cityList.push(city3)
    end
    return cityList
  end

  def make
    @datas.each do |city, collection|
      data = {"type": "FeatureCollection", "features": []}
      data["features"] = collection
      addr = self.checkCity(collection[0]['properties'])
      pref = addr[0]
      city = addr.last
      FileUtils.mkdir_p('geojson/' + pref)
      File.open('geojson/' + pref + '/' + city + '.json', 'w').write(JSON.generate(data))
    end
  end

   def make2
    list = []
    @datas.each do |city, collection|
      addr = self.checkCity(collection[0]['properties'])
      pref = addr[0]
      main = addr.last
      full = addr.join('');
      list.push({'pref': pref, 'city': main, 'full': full})
    end
    File.open('list.json', 'w').write('var list = ' + JSON.generate(list) + ';')
  end
end

city = GeoJsonToCity.new
city.make()
