class CPUStrategyManager
  DIRECTIONS = {
    up: {row: -1, col: 0},
    down: {row: 1, col: 0},
    left: {row: 0, col: -1},
    right: {row: 0, col: 1}
  }.freeze

  OPPOSITE_DIRECTIONS = {
    up: :down,
    down: :up,
    left: :right,
    right: :left
  }.freeze

  def initialize(config)
    @config = config
    @mode = "hunt"
    @target_queue = []
    @last_hit = nil
    @hit_direction = nil
  end

  def next_guess(guesses)
    return process_target_mode(guesses) if @mode == "target" && !@target_queue.empty?
    generate_hunt_guess(guesses)
  end

  def record_hit(row, col, guesses)
    @mode = "target"
    @last_hit = {row: row, col: col}

    if @hit_direction.nil?
      add_adjacent_targets(row, col, guesses)
    else
      add_directional_target(row, col, guesses)
    end
  end

  def record_sunk
    reset_targeting
  end

  private

  def process_target_mode(guesses)
    guess = @target_queue.shift
    return guess unless guesses.include?(guess)

    reset_targeting if @target_queue.empty?
    nil
  end

  def generate_hunt_guess(guesses)
    loop do
      guess = "#{rand(@config.board_size)}#{rand(@config.board_size)}"
      return guess unless guesses.include?(guess)
    end
  end

  def add_adjacent_targets(row, col, guesses)
    DIRECTIONS.each do |direction, offset|
      new_row = row + offset[:row]
      new_col = col + offset[:col]
      add_target_if_valid(new_row, new_col, guesses)
    end
  end

  def add_directional_target(row, col, guesses)
    offset = DIRECTIONS[@hit_direction]
    new_row = row + offset[:row]
    new_col = col + offset[:col]

    if valid_target?(new_row, new_col, guesses)
      @target_queue << "#{new_row}#{new_col}"
    else
      @hit_direction = OPPOSITE_DIRECTIONS[@hit_direction]
      add_directional_target(@last_hit[:row], @last_hit[:col], guesses)
    end
  end

  def add_target_if_valid(row, col, guesses)
    return unless valid_target?(row, col, guesses)
    guess = "#{row}#{col}"
    @target_queue << guess unless @target_queue.include?(guess)
  end

  def valid_target?(row, col, guesses)
    @config.valid_position?(row, col) && !guesses.include?("#{row}#{col}")
  end

  def reset_targeting
    @mode = "hunt"
    @target_queue = []
    @last_hit = nil
    @hit_direction = nil
  end
end
