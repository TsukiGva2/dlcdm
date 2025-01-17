#include <LiquidCrystal_I2C.h>
#include <nanoFORTH.h>

#define LABEL_COUNT 2
#define LIT_COUNT 2

const char* labels[] = {
  "Hello,",
  " World!",
};
const int labels_len[LABEL_COUNT] = {
  6,7
};

const char* literals[] = {
  ":)",
  ":("
};

const char code[] PROGMEM =          ///< define preload Forth code here
": lbl 1 API ;\n"
": val 2 API ;\n"
": fwd 3 API ;\n"
": lit API fwd ;\n";

uint8_t g_x, g_y;

LiquidCrystal_I2C lcd(0x27, 16, 4); // Replace 0x27 with your I2C address

void setup() {
  lcd.init();      // Initialize the LCD
  lcd.backlight(); // Turn on the backlight

  Serial.begin(115200); // On my Arduino Mega
  while(!Serial);

  n4_setup(code);
  n4_api(1, forth_label);
  n4_api(2, forth_literal);
  n4_api(3, forth_fwd);

  // Leave this here if you want a switchScreen button
  pinMode(7, INPUT_PULLUP);
}

void forth_literal() {

  int v = n4_pop();

  if (v > LIT_COUNT || v < 0) {

    return;
  }

  lcd.print(literals[v]);
}

void forth_label() {

  static int current_labels[4] = {-1,-1,-1,-1};

  int v = n4_pop();

  if (v >= LABEL_COUNT || v < 0) {

    return;
  }

  if (v != current_labels[g_y]) {

    forth_clear_line(0);
    current_labels[g_y] = v;
    lcd.print(labels[v]);

  } else {

    forth_clear_line(labels_len[v]);
  }
}

void forth_clear_line(int start) {
  lcd.setCursor(start, g_y);

  for (size_t i = start; i <= 16; i++) {
    lcd.print(' ');
  }

  lcd.setCursor(start, g_y);
}

void forth_fwd() {
  ++g_y %= 4; // number of lines
}

void loop() {
  n4_run();
}

