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
        # 都道府県
        self.setData(city[0], feature)
        # 市町村1
        self.setData(city[0] + city[1], feature)
        # 市町村2
        if !city[2].nil?
          word = city[0] + city[1] + city[2]
          self.setData(word, feature)
        end
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
      data = {'type' => 'FeatureCollection', 'features' => []}
      data[:features] = collection
      addr = self.checkCity(collection[0]['properties'])
      pref = addr[0]
      FileUtils.mkdir_p('geojson/' + pref)
      File.open('geojson/' + pref + '/' + ( ( city == pref ) ? city : city.sub( /^#{pref}/, '' ) ) + '.json', 'w').write(JSON.generate(data))
    end
  end

   def make2
    list = []
    @datas.each do |city, collection|
      p city
      addr = self.checkCity(collection[0]['properties'])
      pref = addr[0]
      if city == pref || city == addr[1] then
        addr1 = ''
        addr2 = ''
      elsif addr[2].nil? then
        addr1 = addr[1]
        addr2 = ''
      else
        addr1 = addr[1]
        addr2 = addr[2]
      end
      list.push({'pref' => pref, 'main' => city, 'addr1' => addr1, 'addr2' => addr2})
    end
    File.open('list.json', 'w').write('var list = ' + JSON.generate(list) + ';')
  end
end

city = GeoJsonToCity.new
city.make()
