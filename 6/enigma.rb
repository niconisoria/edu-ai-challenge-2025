ALPHABET = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

def mod(n, m)
  ((n % m) + m) % m
end

ROTORS = [
  { wiring: 'EKMFLGDQVZNTOWYHXUSPAIBRCJ', notch: 'Q' }, # Rotor I
  { wiring: 'AJDKSIRUXBLHWTMCQGZNPYFVOE', notch: 'E' }, # Rotor II
  { wiring: 'BDFHJLCPRTXVZNYEIWGAKMUSQO', notch: 'V' }  # Rotor III
]

REFLECTOR = 'YRUHQSLDPXNGOKMIEBFZCWVJAT'

def plugboard_swap(c, pairs)
  pairs.each do |a, b|
    return b if c == a
    return a if c == b
  end
  c
end

class Rotor
  attr_accessor :wiring, :notch, :ring_setting, :position

  def initialize(wiring, notch, ring_setting = 0, position = 0)
    @wiring = wiring
    @notch = notch
    @ring_setting = ring_setting
    @position = position
  end

  def step
    @position = mod(@position + 1, 26)
  end

  def at_notch?
    ALPHABET[@position] == @notch
  end

  def forward(c)
    idx = mod(ALPHABET.index(c) + @position - @ring_setting, 26)
    @wiring[idx]
  end

  def backward(c)
    idx = @wiring.index(c)
    ALPHABET[mod(idx - @position + @ring_setting, 26)]
  end
end

class Enigma
  attr_accessor :rotors, :plugboard_pairs

  def initialize(rotor_ids, rotor_positions, ring_settings, plugboard_pairs)
    @rotors = rotor_ids.map.with_index do |id, i|
      Rotor.new(
        ROTORS[id][:wiring],
        ROTORS[id][:notch],
        ring_settings[i],
        rotor_positions[i]
      )
    end
    @plugboard_pairs = plugboard_pairs
  end

  def step_rotors
    # Fix: Check notch positions before stepping to implement proper double-stepping
    # Bug was: checking notch after stepping caused timing issues
    middle_at_notch = @rotors[1].at_notch?
    right_at_notch = @rotors[2].at_notch?
    
    # Step rotors based on notch positions
    @rotors[0].step if middle_at_notch  # Left rotor steps when middle is at notch
    @rotors[1].step if right_at_notch || middle_at_notch  # Middle rotor steps when right is at notch OR when it's at its own notch (double stepping)
    @rotors[2].step  # Right rotor always steps
  end

  def encrypt_char(c)
    return c unless ALPHABET.include?(c)
    
    step_rotors
    c = plugboard_swap(c, @plugboard_pairs)
    
    (@rotors.length - 1).downto(0) do |i|
      c = @rotors[i].forward(c)
    end

    c = REFLECTOR[ALPHABET.index(c)]

    (0...@rotors.length).each do |i|
      c = @rotors[i].backward(c)
    end

    # Fix: Apply plugboard on output for reciprocity
    # Bug was: plugboard only applied on input, breaking Enigma's reciprocal property
    c = plugboard_swap(c, @plugboard_pairs)
    
    c
  end

  def process(text)
    text.upcase.chars.map { |c| encrypt_char(c) }.join
  end
end

def prompt_enigma
  print 'Enter message: '
  message = gets.chomp
  
  print 'Rotor positions (e.g. 0 0 0): '
  rotor_positions = gets.chomp.split.map(&:to_i)
  
  print 'Ring settings (e.g. 0 0 0): '
  ring_settings = gets.chomp.split.map(&:to_i)
  
  print 'Plugboard pairs (e.g. AB CD): '
  plug_str = gets.chomp
  plug_pairs = plug_str.upcase.scan(/([A-Z]{2})/).map { |pair| [pair[0][0], pair[0][1]] }

  enigma = Enigma.new([0, 1, 2], rotor_positions, ring_settings, plug_pairs)
  result = enigma.process(message)
  puts "Output: #{result}"
end

if __FILE__ == $0
  prompt_enigma
end
