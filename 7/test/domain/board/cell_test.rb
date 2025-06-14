require_relative "../../test_helper"

class CellTest < Minitest::Test
  def setup
    super
    @cell = Cell.new
  end

  def test_initial_state
    assert_equal "~", @cell.value
    assert @cell.empty?
    refute @cell.ship?
    refute @cell.hit?
    refute @cell.miss?
  end

  def test_place_ship
    @cell.place_ship
    assert_equal "S", @cell.value
    refute @cell.empty?
    assert @cell.ship?
    refute @cell.hit?
    refute @cell.miss?
  end

  def test_hit
    @cell.hit
    assert_equal "X", @cell.value
    refute @cell.empty?
    refute @cell.ship?
    assert @cell.hit?
    refute @cell.miss?
  end

  def test_miss
    @cell.miss
    assert_equal "O", @cell.value
    refute @cell.empty?
    refute @cell.ship?
    refute @cell.hit?
    assert @cell.miss?
  end

  def test_equality
    cell1 = Cell.new
    cell2 = Cell.new
    cell3 = Cell.new("X")

    assert_equal cell1, cell2
    refute_equal cell1, cell3
  end
end
