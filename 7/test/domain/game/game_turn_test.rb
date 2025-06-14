require_relative "../../test_helper"

class GameTurnTest < Minitest::Test
  def setup
    super
    @player = HumanPlayer.new(@config)
    @opponent = HumanPlayer.new(@config)
    @turn = GameTurn.new(@player, @opponent)
    @opponent.board.set_cell(0, 0, "S")
    
    ship = Ship.new
    ship.add_location("00")
    @opponent.board.ships << ship
  end

  def test_process_guess_hit
    guess = "00"
    assert_output(/HIT/) do
      assert @turn.process_guess(guess)
    end
    assert @player.has_guessed?(guess)
    assert @opponent.board.get_cell(0, 0).hit?
  end

  def test_process_guess_miss
    guess = "11"
    assert_output(/MISS/) do
      assert @turn.process_guess(guess)
    end
    assert @player.has_guessed?(guess)
    assert @opponent.board.get_cell(1, 1).miss?
  end

  def test_process_guess_already_guessed
    guess = "00"
    @player.record_guess(guess)
    assert_output(/already guessed/) do
      refute @turn.process_guess(guess)
    end
  end

  def test_process_guess_already_hit
    guess = "00"
    @opponent.board.ships.first.record_hit(guess)
    assert_output(/already hit/) do
      assert @turn.process_guess(guess)
    end
  end
end
