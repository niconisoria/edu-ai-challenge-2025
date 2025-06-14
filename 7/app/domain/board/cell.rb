class Cell
  EMPTY = "~".freeze
  SHIP = "S".freeze
  HIT = "X".freeze
  MISS = "O".freeze

  attr_reader :value

  def initialize(value = EMPTY)
    @value = value
  end

  def empty?
    @value == EMPTY
  end

  def ship?
    @value == SHIP
  end

  def hit?
    @value == HIT
  end

  def miss?
    @value == MISS
  end

  def place_ship
    @value = SHIP
  end

  def hit
    @value = HIT
  end

  def miss
    @value = MISS
  end

  def to_s
    @value
  end

  def ==(other)
    other.is_a?(Cell) && @value == other.value
  end
end
