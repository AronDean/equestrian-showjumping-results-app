# Equestrian Show Jumping Results Documentation

## Task Description

> The Equestrian Show Jumping Results application has been developed with the primary goal of providing riders in the equestrian show jumping community with a robust platform to manage their performance records and scores. This application aims to streamline the process of recording and tracking show jumping results, making it easier for riders to monitor their progress, evaluate their performance, and compete effectively.

### Purpose
- The purpose of this application is to facilitate the management of show jumping results for riders participating in equestrian competitions.
- It serves as a central repository for recording and storing horse data, jump height information, and ranking details, enabling riders to efficiently log their achievements.

### Goals
- **Efficient Data Management:** The application's core goal is to allow riders to create, update, and maintain their horse data effortlessly. Riders can input essential details about their horses, including age, color, and performance history.
- **Accurate Scoring:** The application provides a mechanism for riders to enter their scores after each competition. Using a predefined multiplier stored in the backend, it calculates scores based on the jump height and rank achieved.
- **Leaderboard:** The application offers a leaderboard feature, allowing riders to see their rankings based on scores obtained over the last 7 days. This competitive aspect fosters engagement and encourages riders to excel in their performances.
- **User-Friendly Interface:** With a user-friendly front-end powered by Vue.js, the application ensures that riders can easily navigate and interact with the platform.
- **Secure User Authentication:** The application prioritizes security by implementing user registration and authentication systems to safeguard user data and ensure that only authorized individuals can access and modify their records using JWT tokens.

In summary, the Equestrian Show Jumping Results application is designed to simplify the process of recording and managing show jumping results for riders. By offering efficient data management, accurate scoring, and a competitive leaderboard, it aims to enhance the overall experience of riders in the equestrian show jumping community.

### Display Overview

**Login Screen:**
- The Login screen is the entry point for registered users. Riders can enter their credentials (email and password) to access their accounts and the application's features.
- ![login-screen](../equestrian-showjumping-results-app/screen-shots/login.png)

**Registration Screen:**
- The Registration screen allows new users to create accounts. Riders can provide their details, including username, email, and password, to register for the application.

- After successfully logging in, riders are presented with a personalized dashboard. This dashboard serves as the central hub for accessing various features and data management tools.
- ![registration-screen](../equestrian-showjumping-results-app/screen-shots/registration.png)

**Leaderboard:**
- After successfully logging in, riders are presented with a Leaderboard screen displays a ranked list of riders based on their scores achieved in the last 7 days. Users can view their own ranking and the rankings of other riders for a competitive perspective.
- ![leaderboard-screen](../equestrian-showjumping-results-app/screen-shots/leaderboard.png)


**Horse Management:**
   - **Create Horse:**
     - The "Create Horse" screen enables users to add new horse details to the database. They can input information such as the horse's name, color, age.
     - ![add-horse-screen](../equestrian-showjumping-results-app/screen-shots/add-horse.png)
   - **Update Horse:**
     - The "Update Horse" screen allows users to edit and update existing horse information.
     - ![update-horse-screen](../equestrian-showjumping-results-app/screen-shots/update-horse.png)
   - **Horse List View:**
     - The "Horse List View" provides a comprehensive list of all the horses associated with a rider's account. Users can quickly access and review their horse profiles.
     - ![horse-list-screen](../equestrian-showjumping-results-app/screen-shots/horse-list.png)

**Score Management:**
   - **Create Score:**
     - The "Create Score" screen permits users to log their scores after each show jumping competition. They can input data including the horse used, jump height, and rank achieved.
     - ![add-game-score-screen](../equestrian-showjumping-results-app/screen-shots/add-score.png)
   - **Score List View:**
     - The "Score List View" offers a detailed list of all recorded scores. Users can filter the list based on specific criteria such as horse, jump height, or rank, making it easy to track their performance.
     - ![score-list-screen](../equestrian-showjumping-results-app/screen-shots/scorelist.png)

## Architecture
- ![flow-chart](../equestrian-showjumping-results-app/screen-shots/design-digram.png)

## Solution Description

### Database
**Database Type**: PostgreSQL
- PostgreSQL, a powerful open-source relational database management system, forms the robust foundation for storing and managing data in the Equestrian Show Jumping Results application.

**Tables**:
1. **Users Table**:
   - Stores user account information, including user ID, username, email, and hashed passwords.

   | Field           | Type         | Description            |
   | --------------- | ------------ | ---------------------- |
   | user_id         | SERIAL       | Unique user identifier |
   | name            | VARCHAR(50)  | User's chosen username |
   | email           | VARCHAR(100) | User's email address   |
   | hashed_password | VARCHAR(100) | Hashed user password   |

2. **Horses Table**:
   - Contains data related to the horses, such as horse ID, name, breed, age, and owner (linked to a user ID).

   | Field    | Type         | Description                  |
   | -------- | ------------ | ---------------------------- |
   | horse_id | SERIAL       | Unique horse identifier      |
   | name     | VARCHAR(100) | Horse's name                 |
   | color    | VARCHAR(50)  | Horse's color                |
   | age      | INTEGER      | Horse's age in years         |
   | owner_id | INTEGER      | User ID of the horse's owner |

3. **Scores Table**:
   - Stores recorded scores for each competition, including score ID, horse ID, jump height ID, rank ID, and the date of the competition.

   | Field            | Type    | Description                           |
   | ---------------- | ------- | ------------------------------------- |
   | score_id         | SERIAL  | Unique score identifier               |
   | horse_id         | INTEGER | Horse ID associated with the score    |
   | jump_height_id   | INTEGER | ID of the Height of the jumps cleared |
   | rank_id          | INTEGER | ID of the rank achieved by the rider  |
   | competition_date | DATE    | Date of the competition               |

4. **Jump Heights Table**:
    - Contains data related to the height of the jumps cleared, including jump height ID and the height in centimeters.

    | Field          | Type        | Description                                               |
    | -------------- | ----------- | --------------------------------------------------------- |
    | jump_height_id | SERIAL      | Unique jump height identifier                             |
    | height         | VARCHAR(50) | Height of the jumps cleared (cm)                          |
    | multiplier     | INTEGER     | Multiplier used to calculate points earned using the rank |

5. **Ranks Table**:
    - Stores data related to the rank achieved by the rider, including rank ID and the rank name.

    | Field   | Type    | Description                     |
    | ------- | ------- | ------------------------------- |
    | rank_id | SERIAL  | Unique rank identifier          |
    | point   | INTEGER | Points associated with the rank |

### Backend
- The backend of the application is built using Go and the Chi framework. It is responsible for handling all the business logic, including data processing, calculations, and storage.

### API

**User Registration Endpoint:**

- **Endpoint:** `/api/v1/user/sign-up`
- **Method:** `POST`
- **Description:** Allows new users to register by providing their name, email, and password.


**User Login Endpoint:**
- **Endpoint:** `/api/v1/user/sign-in`
- **Method:** `POST`
- **Description:** Allows registered users to log in by providing their email and password.

**Refresh Token Endpoint:**
- **Endpoint:** `/api/v1/user/refresh`
- **Method:** `POST`
- **Description:** Allows users to refresh their access token by providing a valid refresh token. This ensures continued access to protected resources without the need for re-authentication.

**Create Horse Endpoint:**
- **Endpoint:** `/api/v1/horse/`
- **Method:** `POST`
- **Description:** Allows users to create a new horse profile by providing details such as the horse's name, color, and age. This information is added to the user's list of horses.

**Update Horse Endpoint:**
- **Endpoint:** `/api/v1/horse/{horse_id}`
- **Method:** `PUT`
- **Description:** Allows users to update an existing horse's details, identified by the `{horse_id}` parameter. Users can modify attributes such as the horse's name, color, and age to keep the information accurate and up to date.

**List Horses Endpoint:**
- **Endpoint:** `/api/v1/horse/`
- **Method:** `GET`
- **Description:** Retrieves a list of all horses associated with the user's account. Users can view and manage their horse profiles conveniently.

**Create Score Endpoint:**
- **Endpoint:** `/api/v1/score/`
- **Method:** `POST`
- **Description:** Allows users to log a new score for a specific competition. Users must provide the `horse_id`, `rank_id`, and `jump_height_id` to record the score accurately. This data is added to the user's list of recorded scores.

**Score List View Endpoint (Filtered):**
- **Endpoint:** `/api/v1/score/?page=1&pageSize=5&horse=1`
- **Method:** `GET`
- **Description:** Retrieves a list of recorded scores, filtered by specific criteria. In this example, scores are filtered by `horse` with ID `1`. Users can also specify pagination with `page` and `pageSize` parameters to manage the number of results displayed per page.

**Leaderboard Endpoint:**
- **Endpoint:** `/api/v1/score/leader-board/`
- **Method:** `GET`
- **Description:** Retrieves a leaderboard displaying the rankings of riders based on their scores achieved in the last 7 days. Users can view the competitive standings within the equestrian community.

### Postman

**Role in the Project:**
- Postman played a crucial role in the development of the Equestrian Show Jumping Results application by serving as a powerful API testing and interaction tool.
- It allowed for comprehensive testing of API endpoints, ensuring that the backend functions correctly and delivers expected responses.
- Postman's flexibility enabled the simulation of various user scenarios, helping to identify and address potential issues early in the development process.

**How It Helped:**
- Postman significantly expedited the testing and validation process, saving valuable development time and resources.
- It provided a user-friendly interface for developers and testers to send requests, inspect responses, and troubleshoot API-related problems efficiently.
- By creating and organizing API requests within Postman collections, the development team achieved a structured and standardized testing approach.
- Postman's environment variables allowed for seamless switching between different testing environments and configurations.
- Overall, Postman's role was instrumental in ensuring the reliability and functionality of the application's API, contributing to a smoother development and testing cycle.

### Authorization and Authentication
- The app employs JWT (JSON Web Tokens) for user authentication.
- Users receive access tokens upon login.
- Refresh tokens are used to renew access tokens, ensuring continuous access.

### Frontend

**Frontend Technologies:**
- The frontend of the Equestrian Show Jumping Results application is developed using Vue.js, with state management powered by Pinia.
   - Pinia provides efficient and reactive state management, ensuring a smooth user experience and robust data handling.
- Vue Router is utilized for seamless navigation within the application.


### Deployment

**Server Requirements:**
- Ensure that you have a server environment meeting the following requirements:
   - Docker installed on the server.
   - Docker Compose installed to orchestrate multiple containers.

**Deployment Process:**
1. **Clone the Repository:**
   - Begin by cloning the repository containing your Equestrian Show Jumping Results application to your server.

2. **Navigate to the Project Directory:**
   - Open a terminal window and navigate to the project directory on your server where you have the Dockerfile and Docker Compose file.

3. **Build Docker Images:**
   - Use the following commands to build Docker images for the frontend and backend (replace `[image_name]` with your desired image names):
     ```bash
     docker build -t [image_name_frontend] ./frontend
     docker build -t [image_name_backend] ./backend
     ```

4. **Compose the Application:**
   - Run the Docker Compose command to create and start containers for the database, frontend, and backend:
     ```bash
     docker-compose up -d
     ```

5. **Access the Application:**
   - Once the containers are up and running, your Equestrian Show Jumping Results application should be accessible. You can access it by navigating to your server's IP address or domain name in a web browser.

**Additional Configuration:**
- Depending on the server environment and specific deployment needs, we may need to configure domain settings, database connections, and environment variables within the Docker Compose file.

### Future Improvements

While the Equestrian Show Jumping Results project has achieved its initial goals, there are several avenues for future enhancements and expansions:

1. **User Profiles:** Implement user profile customization with profile pictures, bio sections, and personal preferences.

2. **Notifications:** Introduce real-time notifications for users to stay informed about score updates, leaderboard changes, and event notifications.

3. **Mobile Application:** Develop a mobile app version to provide on-the-go access and convenience for riders.

4. **Data Analytics:** Incorporate data analytics and visualization tools to offer riders insights into their performance trends.

5. **Internationalization:** Add support for multiple languages to make the application accessible to a broader global audience.

6. **Integration:** Explore integration with third-party equestrian data sources or social media platforms to enrich the user experience.

### Conclusion

The Equestrian Show Jumping Results project is a robust web application designed to streamline score management and foster competition among riders. With a powerful backend, user-friendly frontend, and secure authentication, it offers a comprehensive solution for equestrian enthusiasts. I would like to thank you for exploring this documentation and invite you to reach out with any questions or feedback. Happy riding!
