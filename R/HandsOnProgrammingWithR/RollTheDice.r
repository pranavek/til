roll <- function(die){
  dice <- sample(die,size = 2,replace = TRUE)
  sum(dice)
}

die <- 1:6
roll(die)