import pygame
import random
import math

# settings
WIDTH, HEIGHT = 800, 600
FPS = 60

# colors
BLACK = (0, 0, 0)
WHITE = (255, 255, 255)
RED = (255, 0, 0)
GREEN = (0, 255, 0)
BLUE = (0, 100, 255)
YELLOW = (255, 255, 0)

# player
player = {
    "x": WIDTH // 2,
    "y": HEIGHT - 50,
    "width": 40,
    "height": 20,
    "speed": 10
}

# game state
bullets = []
enemies = []
score = 0
enemy_spawn_timer = 0
game_over = False

pygame.init()
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("Space Shooter")
clock = pygame.time.Clock()
font = pygame.font.Font(None, 36)

running = True
while running:
    # 1. handle input
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

        # shoot on spacebar
        if event.type == pygame.KEYDOWN:
            if event.key == pygame.K_SPACE and not game_over:
                bullets.append({
                    "x": player["x"],
                    "y": player["y"],
                    "vel_y": -8
                })

        # restart on R
        if event.type == pygame.KEYDOWN:
            if event.key == pygame.K_r and game_over:
                bullets = []
                enemies = []
                score = 0
                game_over = False
                player["x"] = WIDTH // 2

    if not game_over:
        # move player with arrow keys
        keys = pygame.key.get_pressed()
        if keys[pygame.K_LEFT] and player["x"] > 20:
            player["x"] -= player["speed"]
        if keys[pygame.K_RIGHT] and player["x"] < WIDTH - 20:
            player["x"] += player["speed"]

        # spawn enemies
        enemy_spawn_timer += 1
        if enemy_spawn_timer >= 60:  # spawn every 60 frames
            enemies.append({
                "x": random.randint(20, WIDTH - 20),
                "y": 0,
                "vel_y": random.uniform(1, 3)
            })
            enemy_spawn_timer = 0

        # update bullets
        for bullet in bullets[:]:
            bullet["y"] += bullet["vel_y"]
            if bullet["y"] < 0:
                bullets.remove(bullet)

        # update enemies
        for enemy in enemies[:]:
            enemy["y"] += enemy["vel_y"]

            # game over if enemy reaches bottom
            if enemy["y"] > HEIGHT:
                game_over = True

        # collision detection — bullet hits enemy
        for bullet in bullets[:]:
            for enemy in enemies[:]:
                distance = math.sqrt(
                    (bullet["x"] - enemy["x"])**2 +
                    (bullet["y"] - enemy["y"])**2
                )
                if distance < 20:
                    bullets.remove(bullet)
                    enemies.remove(enemy)
                    score += 10
                    break

    # 2. draw everything
    screen.fill(BLACK)

    # draw player as triangle
    pygame.draw.polygon(screen, GREEN, [
        (player["x"], player["y"] - 20),       # top
        (player["x"] - 20, player["y"] + 10),  # bottom left
        (player["x"] + 20, player["y"] + 10)   # bottom right
    ])

    # draw bullets
    for bullet in bullets:
        pygame.draw.circle(screen, YELLOW, (int(bullet["x"]), int(bullet["y"])), 4)

    # draw enemies
    for enemy in enemies:
        pygame.draw.circle(screen, RED, (int(enemy["x"]), int(enemy["y"])), 15)

    # draw score
    score_text = font.render(f"Score: {score}", True, WHITE)
    screen.blit(score_text, (10, 10))

    # draw game over
    if game_over:
        go_text = font.render("GAME OVER! Press R to restart", True, RED)
        screen.blit(go_text, (WIDTH//2 - 200, HEIGHT//2))
        final_text = font.render(f"Final Score: {score}", True, WHITE)
        screen.blit(final_text, (WIDTH//2 - 100, HEIGHT//2 + 40))

    pygame.display.flip()
    clock.tick(FPS)

pygame.quit()