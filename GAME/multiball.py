import pygame
import random

# settings
WIDTH, HEIGHT = 700, 600
FPS = 60
GRAVITY = 0.1
BOUNCE = 0.8
ball_radius = 5

# create 100 particles
balls = []
for i in range(1500):
    balls.append({
        "x": random.randint(0, WIDTH),
        "y": random.randint(0, HEIGHT),
        "vel_x": random.uniform(-3, 3),
        "vel_y": random.uniform(-3, 3)
    })

# pygame setup
pygame.init()
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("Particle System")
clock = pygame.time.Clock()

running = True
while running:
    # 1. handle input
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

        # click to spawn particles at mouse position
        if event.type == pygame.MOUSEBUTTONDOWN:
            mouse_x, mouse_y = pygame.mouse.get_pos()
            for ball in balls:
                ball["x"] = mouse_x
                ball["y"] = mouse_y
                ball["vel_x"] = random.uniform(-5, 5)
                ball["vel_y"] = random.uniform(-10, -2)

    # 2. update physics
    for ball in balls:
        ball["vel_y"] += GRAVITY
        ball["x"] += ball["vel_x"]
        ball["y"] += ball["vel_y"]

        # 3. collision detection per ball
        if ball["x"] + ball_radius >= WIDTH or ball["x"] - ball_radius <= 0:
            ball["vel_x"] = -ball["vel_x"]

        if ball["y"] + ball_radius >= HEIGHT:
            ball["y"] = HEIGHT - ball_radius
            ball["vel_y"] = -ball["vel_y"] * BOUNCE

    # 4. draw everything
    screen.fill((0, 0, 0))
    for ball in balls:
        pygame.draw.circle(screen, (255, 0, 0), (int(ball["x"]), int(ball["y"])), ball_radius)
    pygame.display.flip()

    clock.tick(FPS)

pygame.quit()