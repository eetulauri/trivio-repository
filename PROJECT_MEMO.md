# Trivia Game Project Memo

## Project Overview
A web-based trivia game that uses Google's Gemini 2.0 Flash API to generate questions and validate answers. The project aims to create an engaging, simple-to-use trivia game with AI-generated questions.

## Tech Stack
- **Backend**: Go with Fiber framework
- **Frontend**: Vue.js 3 + Tailwind CSS
- **AI**: Google Gemini 2.0 Flash API
- **Future Database**: PostgreSQL
- **Hosting**: Google Cloud Platform (planned)

## Project Phases

### Phase 1: MVP (Current Phase)
- [x] Basic Go backend setup
- [x] Gemini API integration
- [x] Basic API endpoints
- [ ] Simple Vue.js frontend
- [ ] Basic game loop implementation
- [ ] Dark mode UI

### Phase 2: Database Integration
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

## API Endpoints
- `GET /api/health` - Health check
- `GET /api/question` - Get random trivia question
- `POST /api/check-answer` - Validate user's answer

## Current Status
- Backend implementation complete with Gemini API integration
- Frontend development in progress
- Dark mode UI being implemented
- Basic game loop pending implementation

## Development Notes
- Gemini API requires API key (set via GEMINI_API_KEY environment variable)
- Frontend runs on Vite development server
- Backend runs on port 8080 by default
- CORS enabled for local development

## Future Considerations
- Rate limiting for API calls
- Caching strategy for questions
- Error handling improvements
- Mobile responsiveness
- Accessibility features
- Performance monitoring
- Cost optimization for cloud hosting

## Project Goals
1. Create an engaging trivia experience
2. Demonstrate AI integration in web applications
3. Maintain clean, maintainable code structure
4. Ensure scalability for future features
5. Provide responsive and accessible UI
6. Keep infrastructure costs minimal

## Resources
- [Gemini API Documentation](https://ai.google.dev/gemini-api/docs/models/gemini#gemini-2.0-flash)
- [Vue.js Documentation](https://vuejs.org/)
- [Tailwind CSS Documentation](https://tailwindcss.com/)
- [Go Fiber Documentation](https://docs.gofiber.io/) 