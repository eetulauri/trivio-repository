# Trivia Game Project Memo

## Project Overview
A web-based trivia game that uses Google's Gemini Pro API to generate questions and validate answers. The project aims to create an engaging, simple-to-use trivia game with AI-generated questions and educational feedback.

## Tech Stack
- **Backend**: Go with Fiber framework
- **Frontend**: Vue.js 3 + Tailwind CSS
- **AI**: Google Gemini Pro API (official SDK)
- **Future Database**: PostgreSQL
- **Hosting**: Google Cloud Platform (planned)

## Project Phases

### Phase 1: MVP (Current Phase)
- [x] Basic Go backend setup
- [x] Gemini API integration with official SDK
- [x] Basic API endpoints
- [x] Simple Vue.js frontend
- [x] Basic game loop implementation
- [x] Dark mode UI
- [x] Enhanced game mechanics
  - [x] Answer history display (limited to previous question)
  - [x] Educational tidbits
  - [x] Fuzzy answer matching
  - [x] Auto-focus and keyboard shortcuts
  - [x] Improved UI positioning
  - [x] Color-coded feedback
  - [x] Hot reloading setup

### Phase 2: Database Integration (Next Phase)
- [ ] Set up PostgreSQL database
- [ ] Store question history
- [ ] Track user sessions
- [ ] Basic analytics implementation

### Phase 3: Enhancement
- [ ] User authentication
- [ ] Scoring system
- [ ] Categories/difficulty levels
- [ ] Leaderboards
- [ ] Performance optimization

### Phase 4: Deployment
- [ ] Set up Google Cloud infrastructure
- [ ] CI/CD pipeline
- [ ] Monitoring and logging
- [ ] Production deployment

## Current Implementation Details

### Backend
- Using official Gemini Go SDK
- Hot reloading with Air
- Environment variables via .env
- Structured JSON responses for answers
- Enhanced error handling
- CORS enabled for development

### Frontend
- Vue 3 with Composition API
- Tailwind CSS for styling
- Dark mode by default
- Responsive design
- Auto-focusing input field
- Keyboard shortcuts (Enter to submit)
- Transition animations
- Clean, minimalist UI

### Game Flow
1. Question is displayed
2. User inputs answer
3. Answer is validated by AI
4. Feedback and educational tidbit shown
5. Previous question stored in history
6. New question loaded automatically
7. Input field auto-focused

### Development Environment
- Backend runs on port 8080
- Frontend uses Vite dev server
- Hot reloading enabled for both
- Environment variables managed via .env
- Git ignore patterns set up

## API Endpoints
- `GET /api/health` - Health check
- `GET /api/question` - Get random trivia question
- `POST /api/check-answer` - Validate user's answer with feedback

## Development Notes
- Gemini API key required in .env file
- Run backend with `air` for hot reloading
- Run frontend with `npm run dev`
- Frontend auto-reloads on changes
- Backend auto-reloads with Air

## UI/UX Features
- Clean, minimalist design
- Dark mode optimized
- Color-coded feedback (green/red)
- Educational tidbits after each answer
- Previous question history
- Smooth transitions
- Responsive input handling
- Clear visual hierarchy

## Future Considerations
- Rate limiting for API calls
- Caching strategy for questions
- Error handling improvements
- Mobile responsiveness
- Accessibility features
- Performance monitoring
- Cost optimization for cloud hosting
- Question categories
- Difficulty levels
- User accounts
- Leaderboards

## Known Issues/Limitations
- Limited to one previous question in history
- No persistent storage yet
- No user sessions
- No difficulty settings
- No category selection

## Next Steps
1. Implement database integration
2. Add user sessions
3. Implement scoring system
4. Add difficulty levels
5. Add category selection
6. Prepare for cloud deployment

## Resources
- [Gemini API Documentation](https://ai.google.dev/gemini-api/docs/quickstart?lang=go)
- [Vue.js Documentation](https://vuejs.org/)
- [Tailwind CSS Documentation](https://tailwindcss.com/)
- [Go Fiber Documentation](https://docs.gofiber.io/)

## Development Commands
```bash
# Backend
cd backend
air  # for hot reloading

# Frontend
cd frontend
npm run dev
```

## Required Environment Variables
```env
# Backend (.env)
GEMINI_API_KEY=your-api-key-here
PORT=8080

# Future Database Config
DB_HOST=localhost
DB_PORT=5432
DB_NAME=trivia_db
DB_USER=postgres
DB_PASSWORD=your-password-here
``` 