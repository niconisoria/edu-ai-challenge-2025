require_relative "../../test_helper"

class CPUPlayerTest < Minitest::Test
  def setup
    super
    @strategy = CPUStrategyManager.new(@config)
    @player = CPUPlayer.new(@config, @strategy)
  end

  def test_inherits_base_player
    assert_kind_of BasePlayer, @player
  end

  def test_make_guess
    guess = @player.make_guess
    assert_match(/^\d{2}$/, guess)
  end

  def test_record_hit_and_sunk
    assert_respond_to @player, :record_hit
    assert_respond_to @player, :record_sunk
  end
end
