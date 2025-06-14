require_relative "../../test_helper"

class HumanPlayerTest < Minitest::Test
  def setup
    super
    @player = HumanPlayer.new(@config)
  end

  def test_inherits_base_player
    assert_kind_of BasePlayer, @player
  end
end
