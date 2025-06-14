require_relative "base_player"

class HumanPlayer < BasePlayer
  def initialize(config)
    super
  end

  def make_guess(input_processor)
    loop do
      guess = input_processor.get_guess
      return guess if valid_guess?(guess)
      input_processor.show_invalid_guess_message
    end
  end
end
