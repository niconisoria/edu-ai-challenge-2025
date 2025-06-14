require "set"

class Ship
  attr_reader :locations, :hits

  def initialize
    @locations = []
    @hits = Set.new
  end

  def add_location(location)
    @locations << location
  end

  def hit_at?(location)
    @hits.include?(location)
  end

  def record_hit(location)
    return false unless @locations.include?(location)
    return false if @hits.include?(location)
    @hits.add(location)
    true
  end

  def sunk?
    return false if @locations.empty?
    @locations.all? { |loc| @hits.include?(loc) }
  end
end
