import cv2
import mediapipe as mp

def draw_styled_landmarks(frame, holistic):
    # Make detections
    image, results = mediapipe_detection(frame, holistic)

    # Draw face connections
    mp_drawing.draw_landmarks(image, results.face_landmarks, mp_holistic.FACEMESH_TESSELATION, mp_drawing.DrawingSpec(color=(80,110,10), thickness=1, circle_radius=1), mp_drawing.DrawingSpec(color=(80,256,121), thickness=1, circle_radius=1))

    return image
