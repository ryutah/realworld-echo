create domain long_text as text check (length(value) <= 5000);


create domain short_text as text check (length(value) <= 50);


create domain user_id as text check (length(value) <= 256);
