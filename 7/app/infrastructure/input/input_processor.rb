class InputProcessor
  attr_reader :messages

  def initialize(config, messages)
    @config = config
    @messages = messages
  end

  def process_guess(guess)
    return nil unless valid_guess_format?(guess)
    return nil unless valid_guess_position?(guess)
    guess
  end

  def valid_guess_format?(guess)
    if guess.nil? || guess.length != 2
      puts @messages.format_invalid_guess_format.strip
      return false
    end
    true
  end

  def valid_guess_position?(guess)
    row = guess[0].to_i
    col = guess[1].to_i

    if !@config.valid_position?(row, col)
      puts @messages.format_invalid_guess_position.strip
      return false
    end
    true
  end

  def parse_guess(guess)
    return nil unless guess
    {
      row: guess[0].to_i,
      col: guess[1].to_i,
      string: guess
    }
  end
end
