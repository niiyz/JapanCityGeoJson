require 'json'
require 'fileutils'

class GeoJsonToCity

  def initialize
    @check = {}
    @data  = {}
  end

  def get_info(path)
    Dir.glob(path).each do |path|
      str      = File.open(path).read(nil, '')
      json     = JSON.parse(str)
      features = json['features']
      features.each do |feature|
        # 名称抽出
        addr = self.parse_properties(feature['properties'])
        # properties削除
        feature.delete('properties')
        unless @check.has_key?(addr[:code])
          if @data[addr[:pref_code]].nil?
            @data[addr[:pref_code]] = []
          end
          @data[addr[:pref_code]].push(addr)
          @check[addr[:code]] = 1
        end
      end
    end
  end

  def parse_properties(prop)
    pref     = prop['N03_001']
    city1    = prop['N03_002']
    city2    = prop['N03_003']
    city3    = prop['N03_004']
    code     = prop['N03_007']

    city = ''
    unless city1.nil? && pref != '北海道'
      city += city1
    end
    unless city2.nil?
      city += city2
    end
    unless city3.nil?
      city += city3
    end
    if code.nil?
      # 未所属の場合、県名から都道府県コード取得
      pref_code = self.get_pref_code(pref)
    else
      # 行政区分コード先頭2文字が都道府県コード
      pref_code = code[0, 2]
    end
    {:pref => pref, :city => city, :pref_code => pref_code, :code => code}
  end

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

  def make_read_me(path)
    file = File.open(path, 'w')
    file.puts "|  都道府県  | 都道府県コード | 行政区分 | 行政区分コード |"
    file.puts "|-----------|--------------|--------- |--------------|"
    file
  end

  def make

    @data.each do |key,collection|

      geo_readme  = self.make_read_me("geojson/#{key}/README.md")
      topo_readme = self.make_read_me("topojson/#{key}/README.md")

      collection.each do |info|
        line = "| #{info[:pref]} | #{info[:pref_code]} | #{info[:city]} | #{info[:code]} |"
        geo_readme.puts line
        topo_readme.puts line
      end

      geo_readme.close
      topo_readme.close

    end

  end

end

city = GeoJsonToCity.new
japan_json = './data/geojson/japan2016.json'
city.get_info(japan_json)

city.make()

