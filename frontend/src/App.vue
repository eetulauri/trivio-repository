<template>
  <div class="min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100">
    <div class="container mx-auto px-4 py-4">
      <header class="text-center mb-6">
        <h1 class="text-4xl font-bold text-primary-600 dark:text-primary-400">Trivia Game</h1>
        <p class="text-gray-600 dark:text-gray-400 mt-2">Test your knowledge with AI-generated questions!</p>
      </header>

      <main class="max-w-2xl mx-auto space-y-6">
        <!-- Previous Question Section -->
        <div v-if="history.length > 0" class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6">
          <div class="space-y-4">
            <div class="border-b border-gray-200 dark:border-gray-700 pb-4">
              <h3 class="text-lg font-medium mb-2">Previous Question:</h3>
              <p class="mb-2">{{ history[0].question }}</p>
              <p class="text-sm text-gray-600 dark:text-gray-400">Your answer: {{ history[0].answer }}</p>
              <div :class="[
                'mt-3 p-3 rounded',
                history[0].correct ? 'bg-green-100 dark:bg-green-900 text-green-700 dark:text-green-100' : 'bg-red-100 dark:bg-red-900 text-red-700 dark:text-red-100'
              ]">
                <p class="font-medium">{{ history[0].feedback }}</p>
                <p class="mt-2 text-sm">{{ history[0].tidbit }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Current Question Section -->
        <div v-if="loading" class="text-center">
          <div class="animate-spin rounded-full h-12 w-12 border-4 border-primary-500 border-t-transparent mx-auto"></div>
          <p class="mt-4 text-gray-600 dark:text-gray-400">Loading question...</p>
        </div>

        <div v-else class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6">
          <h2 class="text-xl font-semibold mb-4">{{ currentQuestion || 'Ready to start?' }}</h2>
          
          <form v-if="currentQuestion" @submit.prevent="checkAnswer" class="space-y-4">
            <div>
              <input 
                ref="answerInput"
                v-model="userAnswer"
                type="text"
                placeholder="Type your answer..."
                class="w-full px-6 py-3 rounded-lg border dark:border-gray-700 bg-gray-50 dark:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-primary-500 text-lg"
                :disabled="checking"
                @keyup.enter="checkAnswer"
              />
            </div>
            
            <button
              type="submit"
              class="w-full bg-primary-500 hover:bg-primary-600 text-white font-semibold py-3 px-6 rounded-lg transition-colors text-lg"
              :disabled="!userAnswer || checking"
            >
              Submit Answer
            </button>
          </form>

          <button
            v-else
            @click="getNewQuestion"
            class="w-full bg-primary-500 hover:bg-primary-600 text-white font-semibold py-3 px-6 rounded-lg transition-colors text-lg"
          >
            Start Game
          </button>
        </div>

        <!-- Keyboard Shortcuts Help -->
        <div class="text-center text-sm text-gray-600 dark:text-gray-400">
          <p>Press <kbd class="px-2 py-1 bg-gray-200 dark:bg-gray-700 rounded">Enter</kbd> to submit answer</p>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick, watch } from 'vue'

const API_BASE_URL = 'http://localhost:8080/api'

const currentQuestion = ref('')
const userAnswer = ref('')
const loading = ref(false)
const checking = ref(false)
const history = ref([])
const answerInput = ref(null)

// Keep history limited to one item
watch(history, (newHistory) => {
  if (newHistory.length > 1) {
    history.value = [newHistory[0]]
  }
})

// Focus input whenever it's available
watch(currentQuestion, async () => {
  await nextTick()
  if (answerInput.value) {
    answerInput.value.focus()
  }
})

const getNewQuestion = async () => {
  loading.value = true
  userAnswer.value = ''
  
  try {
    const response = await fetch(`${API_BASE_URL}/question`)
    const data = await response.json()
    currentQuestion.value = data.question
  } catch (error) {
    console.error('Error fetching question:', error)
    currentQuestion.value = 'Error loading question. Please try again.'
  } finally {
    loading.value = false
  }
}

const checkAnswer = async () => {
  if (!userAnswer.value || checking.value) return
  
  checking.value = true
  try {
    const response = await fetch(`${API_BASE_URL}/check-answer`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        question: currentQuestion.value,
        answer: userAnswer.value,
      }),
    })
    
    const data = await response.json()
    
    // Add to history
    history.value.unshift({
      question: currentQuestion.value,
      answer: userAnswer.value,
      correct: data.correct,
      feedback: data.feedback,
      tidbit: data.tidbit
    })

    // Get next question automatically
    await getNewQuestion()
  } catch (error) {
    console.error('Error checking answer:', error)
  } finally {
    checking.value = false
  }
}

// Enable dark mode by default
onMounted(() => {
  document.documentElement.classList.add('dark')
  getNewQuestion() // Start the game automatically
})
</script>

<style>
.container {
  max-width: 1200px;
}

/* Smooth transitions */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
