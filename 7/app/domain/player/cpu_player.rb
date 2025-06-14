require_relative "base_player"

class CPUPlayer < BasePlayer
  def initialize(config, strategy)
    super(config)
    @strategy = strategy
  end

  def make_guess
    @strategy.next_guess(@guesses)
  end

  def record_hit(row, col)
    @strategy.record_hit(row, col, @guesses)
  end

  def record_sunk
    @strategy.record_sunk
  end
end
