require_relative "../../test_helper"

class ShipTest < Minitest::Test
  def setup
    super
    @ship = Ship.new
  end

  def test_initial_state
    assert_empty @ship.locations
    assert_empty @ship.hits
    refute @ship.sunk?
  end

  def test_add_location
    @ship.add_location("00")
    assert_includes @ship.locations, "00"
  end

  def test_hit_at
    @ship.add_location("00")
    refute @ship.hit_at?("00")

    @ship.record_hit("00")
    assert @ship.hit_at?("00")
  end

  def test_record_hit
    @ship.add_location("00")
    assert @ship.record_hit("00")
    assert_includes @ship.hits, "00"

    # Can't hit the same location twice
    refute @ship.record_hit("00")
  end

  def test_sunk
    @ship.add_location("00")
    @ship.add_location("01")

    refute @ship.sunk?

    @ship.record_hit("00")
    refute @ship.sunk?

    @ship.record_hit("01")
    assert @ship.sunk?
  end

  def test_invalid_hit
    @ship.add_location("00")
    refute @ship.record_hit("11") # Hit at non-ship location
    assert_empty @ship.hits
  end
end
