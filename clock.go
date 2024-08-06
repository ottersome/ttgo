package main

var BIG_NUMS = [10][15]int{
  {1,1,1,1,0,1,1,0,1,1,0,1,1,1,1}, /* 0 */
  {0,0,1,0,0,1,0,0,1,0,0,1,0,0,1}, /* 1 */
  {1,1,1,0,0,1,1,1,1,1,0,0,1,1,1}, /* 2 */
  {1,1,1,0,0,1,1,1,1,0,0,1,1,1,1}, /* 3 */
  {1,0,1,1,0,1,1,1,1,0,0,1,0,0,1}, /* 4 */
  {1,1,1,1,0,0,1,1,1,0,0,1,1,1,1}, /* 5 */
  {1,1,1,1,0,0,1,1,1,1,0,1,1,1,1}, /* 6 */
  {1,1,1,0,0,1,0,0,1,0,0,1,0,0,1}, /* 7 */
  {1,1,1,1,0,1,1,1,1,1,0,1,1,1,1}, /* 8 */
  {1,1,1,1,0,1,1,1,1,0,0,1,1,1,1}, /* 9 */
}
//For Space between digits with colon
var BIG_SPACE = [15]int{
  0,0,0,0,1,0,0,0,0,0,1,0,0,0,0,
}

type clock struct {
  hour int
  minute int
  seconds int
}

//CHECK: We are assumign utf-8 for now.
//Might want to add support for other encodings later.
var fg_character = []rune{0x2588, 0x2584, 0x2582, 0x2583}

// Dont know, we might want to add more customization rather than just drawing it
func GetNumRender(num int) [30]rune {
  num_array := BIG_NUMS[num]
  rune_buffer := RenderRune(num_array)
  return rune_buffer
}
    
func RenderRune(num_array [15]int) [30]rune{
  string_buffer := [30]rune{}
  for i := 0; i < 30; i++ {
    char_to_use := fg_character[num_array[i/2]]
    // Make sure this UTF-8 int can 
    string_buffer[i] = char_to_use
  }
  return string_buffer
}
func (c clock) get_string() string {
  final_time_str := []rune{}
  // We Need to convert 4 digits total. Two from hour, two from minute
  
  //Hour digits:
  hour_digit_left := GetNumRender(c.hour/10)
  hour_digit_right := GetNumRender(c.hour%10)

  // Space In Between
  in_between := RenderRune(BIG_SPACE)

  //Minute digits:
  minute_digit_left := GetNumRender(c.minute/10)
  minute_digit_right := GetNumRender(c.minute%10)

  // Seconds digits: (for later)
  // second_digit_left := m.currenct_clock.GetString(m.currenct_clock.seconds/10)
  // second_digit_right := m.currenct_clock.GetString(m.currenct_clock.seconds%10)

  // All digits to be consdiered
  tbc := [5][30]rune{
    hour_digit_left,
    hour_digit_right,
    in_between,
    minute_digit_left,
    minute_digit_right,
  }
  //offsets will be a map that will map tbc indices to x,y offsets
  // offsets := map[int][2]int{
  //   0: [2]int{0,0},
  //   1: [2]int{0,0},
  //   2: [2]int{0,0},
  //   3: [2]int{0,0},
  //   4: [2]int{0,0},
  // }

  // For each big num we add it one by one:
  for i := 0; i < 5; i++ {
    for j := 0; j < 30; j++ {
      final_time_str = append(final_time_str, tbc[i][j])
    }
  }
  return string(final_time_str[:])
}
