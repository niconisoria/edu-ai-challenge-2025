class Position
  attr_reader :row, :col

  def initialize(row, col)
    @row = row
    @col = col
  end

  def to_s
    "#{@row}#{@col}"
  end

  def self.from_string(str)
    return nil unless str.match?(/^\d{2}$/)
    new(str[0].to_i, str[1].to_i)
  end

  def valid?(board_size)
    @row.between?(0, board_size - 1) && @col.between?(0, board_size - 1)
  end

  def ==(other)
    other.is_a?(Position) && @row == other.row && @col == other.col
  end

  def eql?(other)
    self == other
  end

  def hash
    [@row, @col].hash
  end
end
