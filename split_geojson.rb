require 'json'
require 'fileutils'

class GeoJsonToCity

  def initialize()
    @datas = {}
    @data_info = {}
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
      # data_info
      unless @data_info.has_key?(key)
        @data_info[key] = {:dir_name => dir_name, :file_name => file_name}
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
  def get_pref_code(pref_name)
    pref = %w[
        北海道 青森県 岩手県 宮城県 秋田県 山形県 福島県 茨城県 栃木県 群馬県 埼玉県
        千葉県 東京都 神奈川県 新潟県 富山県 石川県 福井県 山梨県 長野県 岐阜県 静岡県
        愛知県 三重県 滋賀県 京都府 大阪府 兵庫県 奈良県 和歌山県 鳥取県 島根県 岡山県
        広島県 山口県 徳島県 香川県 愛媛県 高知県 福岡県 佐賀県 長崎県 熊本県 大分県
        宮崎県 鹿児島県 沖縄県
    ]

    idx = pref.index(pref_name)
    sprintf('%02d', idx + 1)
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
