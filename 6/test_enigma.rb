require 'simplecov'
SimpleCov.start do
  add_filter '/test/'
end

require 'minitest/autorun'
require 'minitest/reporters'
require_relative 'enigma'

# Use spec-style output for better readability
Minitest::Reporters.use! Minitest::Reporters::SpecReporter.new

class TestEnigma < Minitest::Test
  
  def test_mod_function
    # Test positive modulo
    assert_equal 1, mod(7, 3)
    assert_equal 0, mod(6, 3)
    
    # Test negative modulo (Ruby's % can return negative results)
    assert_equal 2, mod(-1, 3)
    assert_equal 1, mod(-2, 3)
    assert_equal 0, mod(-3, 3)
    
    # Test with 26 (alphabet size)
    assert_equal 0, mod(26, 26)
    assert_equal 25, mod(-1, 26)
    assert_equal 1, mod(27, 26)
  end

  def test_plugboard_swap
    pairs = [['A', 'B'], ['C', 'D']]
    
    # Test swapping
    assert_equal 'B', plugboard_swap('A', pairs)
    assert_equal 'A', plugboard_swap('B', pairs)
    assert_equal 'D', plugboard_swap('C', pairs)
    assert_equal 'C', plugboard_swap('D', pairs)
    
    # Test no swap for unmapped letters
    assert_equal 'E', plugboard_swap('E', pairs)
    assert_equal 'Z', plugboard_swap('Z', pairs)
    
    # Test empty pairs
    assert_equal 'A', plugboard_swap('A', [])
  end

  def test_rotor_initialization
    rotor = Rotor.new('EKMFLGDQVZNTOWYHXUSPAIBRCJ', 'Q', 5, 10)
    
    assert_equal 'EKMFLGDQVZNTOWYHXUSPAIBRCJ', rotor.wiring
    assert_equal 'Q', rotor.notch
    assert_equal 5, rotor.ring_setting
    assert_equal 10, rotor.position
  end

  def test_rotor_step
    rotor = Rotor.new('EKMFLGDQVZNTOWYHXUSPAIBRCJ', 'Q', 0, 0)
    
    # Test stepping
    rotor.step
    assert_equal 1, rotor.position
    
    # Test wrapping around
    rotor.position = 25
    rotor.step
    assert_equal 0, rotor.position
  end

  def test_rotor_at_notch
    rotor = Rotor.new('EKMFLGDQVZNTOWYHXUSPAIBRCJ', 'Q', 0, 0)
    
    # Q is at position 16 in alphabet
    rotor.position = 16
    assert rotor.at_notch?
    
    rotor.position = 0
    refute rotor.at_notch?
    
    rotor.position = 15
    refute rotor.at_notch?
  end

  def test_rotor_forward_transformation
    rotor = Rotor.new('EKMFLGDQVZNTOWYHXUSPAIBRCJ', 'Q', 0, 0)
    
    # Test with position 0 and ring setting 0
    # A (position 0) should map to E (first letter in wiring)
    assert_equal 'E', rotor.forward('A')
    
    # Test with different position
    rotor.position = 1
    # With position 1, A becomes B in the rotor, which maps to K
    assert_equal 'K', rotor.forward('A')
  end

  def test_rotor_backward_transformation
    rotor = Rotor.new('EKMFLGDQVZNTOWYHXUSPAIBRCJ', 'Q', 0, 0)
    
    # Test with position 0 and ring setting 0
    # E is at position 0 in wiring, so it should map back to A
    assert_equal 'A', rotor.backward('E')
    
    # K is at position 1 in wiring, so it should map back to B
    assert_equal 'B', rotor.backward('K')
  end

  def test_enigma_initialization
    enigma = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [['A', 'B']])
    
    assert_equal 3, enigma.rotors.length
    assert_equal [['A', 'B']], enigma.plugboard_pairs
    
    # Check that rotors are properly initialized with correct wiring
    assert_equal ROTORS[0][:wiring], enigma.rotors[0].wiring
    assert_equal ROTORS[1][:wiring], enigma.rotors[1].wiring
    assert_equal ROTORS[2][:wiring], enigma.rotors[2].wiring
  end

  def test_enigma_step_rotors_normal
    enigma = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
    
    # Normal stepping - only rightmost rotor steps
    initial_positions = enigma.rotors.map(&:position)
    enigma.step_rotors
    
    assert_equal initial_positions[0], enigma.rotors[0].position  # Left rotor unchanged
    assert_equal initial_positions[1], enigma.rotors[1].position  # Middle rotor unchanged
    assert_equal (initial_positions[2] + 1) % 26, enigma.rotors[2].position  # Right rotor stepped
  end

  def test_enigma_step_rotors_double_stepping
    enigma = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
    
    # Set middle rotor at its notch (E for rotor II)
    enigma.rotors[1].position = ALPHABET.index('E')  # Position 4
    
    initial_left = enigma.rotors[0].position
    initial_middle = enigma.rotors[1].position
    enigma.step_rotors
    
    # Both left and middle rotors should step when middle rotor is at notch (double stepping)
    assert_equal (initial_left + 1) % 26, enigma.rotors[0].position
    assert_equal (initial_middle + 1) % 26, enigma.rotors[1].position
  end

  def test_enigma_encrypt_char_non_alphabetic
    enigma = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
    
    # Non-alphabetic characters should pass through unchanged
    assert_equal '1', enigma.encrypt_char('1')
    assert_equal ' ', enigma.encrypt_char(' ')
    assert_equal '!', enigma.encrypt_char('!')
  end

  def test_enigma_encrypt_char_alphabetic
    enigma = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
    
    # Test that encryption produces a valid alphabetic character
    result = enigma.encrypt_char('A')
    assert_includes ALPHABET, result
    refute_equal 'A', result  # Should be different from input
  end

  def test_enigma_encrypt_char_with_plugboard
    # Test plugboard functionality by checking that swapped letters behave correctly
    enigma_with_plug = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [['A', 'B']])
    enigma_without_plug = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
    
    # Encrypt 'A' with plugboard - should behave like encrypting 'B' without plugboard
    result_a_with_plug = enigma_with_plug.encrypt_char('A')
    result_b_without_plug = enigma_without_plug.encrypt_char('B')
    
    # Due to reciprocal plugboard application, 'A' with plugboard should equal 'B' without plugboard
    assert_equal result_a_with_plug, result_b_without_plug
  end

  def test_enigma_process_text
    enigma = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
    
    # Test processing a simple message
    result = enigma.process('HELLO')
    assert_equal 5, result.length
    assert result.chars.all? { |c| ALPHABET.include?(c) }
    
    # Test that it handles lowercase
    enigma2 = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
    result2 = enigma2.process('hello')
    assert_equal result, result2  # Should be same as uppercase
  end

  def test_enigma_reciprocal_property_without_plugboard
    # Test that Enigma is reciprocal without plugboard
    settings = [[0, 1, 2], [5, 10, 15], [0, 0, 0], []]
    
    enigma1 = Enigma.new(*settings)
    enigma2 = Enigma.new(*settings)
    
    original = 'HELLO'
    encrypted = enigma1.process(original)
    decrypted = enigma2.process(encrypted)
    
    assert_equal original, decrypted
  end

  def test_enigma_reciprocal_property_with_plugboard
    # Test that Enigma is reciprocal with plugboard (encoding twice should return original)
    settings = [[0, 1, 2], [5, 10, 15], [0, 0, 0], [['A', 'B'], ['C', 'D']]]
    
    enigma1 = Enigma.new(*settings)
    enigma2 = Enigma.new(*settings)
    
    original = 'HELLO'
    encrypted = enigma1.process(original)
    decrypted = enigma2.process(encrypted)
    
    assert_equal original, decrypted
  end

  def test_enigma_different_rotor_positions
    # Test that different rotor positions produce different results
    enigma1 = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
    enigma2 = Enigma.new([0, 1, 2], [1, 0, 0], [0, 0, 0], [])
    
    result1 = enigma1.process('A')
    result2 = enigma2.process('A')
    
    refute_equal result1, result2
  end

  def test_enigma_different_ring_settings
    # Test that different ring settings produce different results
    enigma1 = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
    enigma2 = Enigma.new([0, 1, 2], [0, 0, 0], [1, 0, 0], [])
    
    result1 = enigma1.process('A')
    result2 = enigma2.process('A')
    
    refute_equal result1, result2
  end

  def test_rotor_constants
    # Test that rotor constants are properly defined
    assert_equal 3, ROTORS.length
    
    ROTORS.each do |rotor|
      assert_equal 26, rotor[:wiring].length
      assert_includes ALPHABET, rotor[:notch]
      assert_equal 26, rotor[:wiring].chars.uniq.length  # All letters unique
    end
    
    # Test reflector
    assert_equal 26, REFLECTOR.length
    assert_equal 26, REFLECTOR.chars.uniq.length  # All letters unique
  end

  def test_alphabet_constant
    assert_equal 26, ALPHABET.length
    assert_equal 'ABCDEFGHIJKLMNOPQRSTUVWXYZ', ALPHABET
  end

  def test_rotor_wiring_permutations
    # Test that each rotor wiring is a valid permutation of the alphabet
    ROTORS.each_with_index do |rotor, index|
      wiring_chars = rotor[:wiring].chars.sort
      alphabet_chars = ALPHABET.chars.sort
      assert_equal alphabet_chars, wiring_chars, "Rotor #{index + 1} should be a permutation of alphabet"
    end
  end

  def test_reflector_is_involution
    # Test that reflector is an involution (applying it twice returns original)
    ALPHABET.each_char do |char|
      reflected_once = REFLECTOR[ALPHABET.index(char)]
      reflected_twice = REFLECTOR[ALPHABET.index(reflected_once)]
      assert_equal char, reflected_twice, "Reflector should be an involution for #{char}"
    end
  end

  def test_enigma_never_encrypts_to_itself
    # Test that no character encrypts to itself (historical Enigma property)
    enigma = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
    
    ALPHABET.each_char do |char|
      # Create fresh enigma for each test to avoid rotor advancement
      fresh_enigma = Enigma.new([0, 1, 2], [0, 0, 0], [0, 0, 0], [])
      result = fresh_enigma.encrypt_char(char)
      refute_equal char, result, "#{char} should not encrypt to itself"
    end
  end
end
