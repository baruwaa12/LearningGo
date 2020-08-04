# Computing-Project - Name TBD (Flappy Bird Clone)

This my new attempt at making my flappy bird clone in golang.

There are currently 6 main tasks which need to be developed for me to complete this game.
  1. Scrolling Background.
  2. Bird Flight
  3. Obstacle Placement
  4. Collison
  5. Death
  6. Scoring
  7. Addition of Levels (Extra)
  8. Creation of a character.
  

First stage of development - Scrolling background.
  
Since I am creating a flappy bird clone. The first thing I noticed with the game was the fact that it had a parallax effect. I collected 2 images. The skyline where you can see the buildings in the distance and the actual foreground.
The skyline background would move the slowest to give the player the impression that there is considerable distance between skyline and foreground. Then the foreground background would move the fastest to show that player is moving at a reasonable speed over the ground
 
 
Second stage of development - Animation of sprites
 
To show movement of the character I need sprites. Sprites are individual images which was layered they convey the impression of movement similar to a filpbook.
 As a testing each stage of development individually I collected a 5 images of a characters walking sprite and programmed them to cycle through every image after every 1/4 of second. I chose this time so that the images can flow nicely without any lag.

Third and fourth stage of development - Obstacle placement

Obstacles will be placed in the players ways for them to avoid. They will need to be programmed so that if the player makes contact with any of the pixels of the obstacle the game should end. They should also have completely random placement on the screen.

Fifth stage of development - Death

Death will occur either when the player touchs the bottom or top of the screen or when they come into contact with any of the obstacle placements.

Sixth stage of the development - Scoring.

There will be 2 score counters. The ring counter and pipe counter. For each ring collected and pipe passed each respective counter should increase by 1. At the end of the game the number of pipes passed and rings collected should be displayed.




Todo - Figure out how to use the scrolling background.