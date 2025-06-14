require_relative "../../test_helper"

class PositionTest < Minitest::Test
  def test_initialize_and_to_s
    pos = Position.new(2, 3)
    assert_equal 2, pos.row
    assert_equal 3, pos.col
    assert_equal "23", pos.to_s
  end

  def test_from_string
    pos = Position.from_string("45")
    assert_instance_of Position, pos
    assert_equal 4, pos.row
    assert_equal 5, pos.col
    assert_nil Position.from_string("a5")
    assert_nil Position.from_string("123")
  end

  def test_valid
    pos = Position.new(1, 1)
    assert pos.valid?(3)
    refute Position.new(-1, 0).valid?(3)
    refute Position.new(0, 3).valid?(3)
  end

  def test_equality
    pos1 = Position.new(1, 2)
    pos2 = Position.new(1, 2)
    pos3 = Position.new(2, 1)
    assert_equal pos1, pos2
    refute_equal pos1, pos3
    assert_equal pos1.hash, pos2.hash
  end
end
