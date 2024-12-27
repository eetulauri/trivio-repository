<template>
  <div class="min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100">
    <div class="container mx-auto px-4 py-8">
      <header class="text-center mb-8">
        <h1 class="text-4xl font-bold text-primary-600 dark:text-primary-400">Trivia Game</h1>
        <p class="text-gray-600 dark:text-gray-400 mt-2">Test your knowledge with AI-generated questions!</p>
      </header>

      <main class="max-w-2xl mx-auto">
        <div v-if="loading" class="text-center">
          <div class="animate-spin rounded-full h-12 w-12 border-4 border-primary-500 border-t-transparent mx-auto"></div>
          <p class="mt-4 text-gray-600 dark:text-gray-400">Loading question...</p>
        </div>

        <div v-else>
          <div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6 mb-6">
            <h2 class="text-xl font-semibold mb-4">{{ currentQuestion || 'Ready to start?' }}</h2>
            
            <form v-if="currentQuestion" @submit.prevent="checkAnswer" class="space-y-4">
              <div>
                <input 
                  v-model="userAnswer"
                  type="text"
                  placeholder="Type your answer..."
                  class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-gray-50 dark:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-primary-500"
                  :disabled="checking"
                />
              </div>
              
              <button
                type="submit"
                class="w-full bg-primary-500 hover:bg-primary-600 text-white font-semibold py-2 px-4 rounded-lg transition-colors"
                :disabled="!userAnswer || checking"
              >
                Submit Answer
              </button>
            </form>

            <button
              v-else
              @click="getNewQuestion"
              class="w-full bg-primary-500 hover:bg-primary-600 text-white font-semibold py-2 px-4 rounded-lg transition-colors"
            >
              Start Game
            </button>
          </div>

          <div v-if="result !== null" :class="[
            'p-4 rounded-lg mb-6 text-center',
            result ? 'bg-green-100 dark:bg-green-900 text-green-700 dark:text-green-100' : 'bg-red-100 dark:bg-red-900 text-red-700 dark:text-red-100'
          ]">
            <p class="font-semibold">{{ result ? 'Correct!' : 'Incorrect!' }}</p>
            <button
              @click="getNewQuestion"
              class="mt-4 text-sm underline hover:no-underline"
            >
              Next Question
            </button>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const API_BASE_URL = 'http://localhost:8080/api'

const currentQuestion = ref('')
const userAnswer = ref('')
const loading = ref(false)
const checking = ref(false)
const result = ref(null)

const getNewQuestion = async () => {
  loading.value = true
  result.value = null
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
  if (!userAnswer.value) return
  
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
    result.value = data.correct
  } catch (error) {
    console.error('Error checking answer:', error)
  } finally {
    checking.value = false
  }
}

// Enable dark mode by default
onMounted(() => {
  document.documentElement.classList.add('dark')
})
</script>

<style>
/* Add any custom styles here */
</style>
