require "set"
require_relative "../board/board"
require_relative "../board/position"

class BasePlayer
  attr_reader :board, :num_ships, :guesses

  def initialize(config)
    @config = config
    @board = Board.new(config)
    @num_ships = config.num_ships
    @guesses = Set.new
  end

  def place_ships
    @board.place_ships_randomly(@num_ships)
  end

  def record_guess(guess)
    @guesses.add(guess)
  end

  def has_guessed?(guess)
    @guesses.include?(guess)
  end

  def decrement_ships
    @num_ships -= 1
  end

  def all_ships_sunk?
    @num_ships.zero?
  end

  def remaining_ships
    @num_ships
  end

  def valid_guess?(guess)
    return false if has_guessed?(guess)
    position = Position.from_string(guess)
    position&.valid?(@config.board_size)
  end
end
