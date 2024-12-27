<template>
  <div class="min-h-screen bg-gray-900 text-gray-100">
    <div class="container max-w-3xl px-4 mt-8">
      <header class="text-center mb-12">
        <h1 class="text-4xl font-bold bg-gradient-to-r from-blue-500 to-purple-500 bg-clip-text text-transparent">
          Trivia Game
        </h1>
      </header>

      <main class="space-y-8">
        <!-- Current Question Section -->
        <section class="bg-gray-800/50 backdrop-blur rounded-xl shadow-xl p-6 border-2 border-gray-700">
          <h2 class="text-lg text-gray-400 uppercase tracking-wide mb-4">Current Question</h2>
          
          <div v-if="loading" class="text-center py-8">
            <div class="animate-spin rounded-full h-12 w-12 border-4 border-blue-500 border-t-transparent mx-auto"></div>
            <p class="mt-4 text-gray-400">Loading question...</p>
          </div>

          <div v-else>
            <p class="text-2xl font-medium mb-6 text-center text-white">
              {{ currentQuestion || 'Ready to start?' }}
            </p>
            
            <form v-if="currentQuestion" @submit.prevent="checkAnswer" class="flex gap-4 items-center">
              <input 
                ref="answerInput"
                v-model="userAnswer"
                type="text"
                placeholder="Type your answer..."
                class="flex-1 px-6 py-4 rounded-xl border-2 border-gray-700 bg-gray-900/50 focus:outline-none focus:ring-2 focus:ring-blue-500 text-xl text-center placeholder:text-gray-500"
                :disabled="checking"
                @keyup.enter="checkAnswer"
              />
              
              <button
                type="submit"
                class="px-8 py-4 bg-gradient-to-r from-blue-500 to-purple-500 hover:from-blue-600 hover:to-purple-600 text-white font-semibold rounded-xl transition-colors text-xl flex items-center justify-center gap-2 whitespace-nowrap"
                :disabled="!userAnswer || checking"
              >
                <span v-if="checking" class="animate-spin rounded-full h-5 w-5 border-2 border-white border-t-transparent"></span>
                {{ checking ? 'Checking...' : 'Submit' }}
              </button>
            </form>

            <button
              v-else
              @click="getNewQuestion"
              class="w-full bg-gradient-to-r from-blue-500 to-purple-500 hover:from-blue-600 hover:to-purple-600 text-white font-semibold py-4 px-8 rounded-xl transition-colors text-xl"
            >
              Start Game
            </button>
          </div>
        </section>

        <!-- Previous Question Section -->
        <section v-if="history.length > 0" class="bg-gray-800/50 backdrop-blur rounded-xl shadow-xl p-6 border-2 border-gray-700">
          <h2 class="text-lg text-gray-400 uppercase tracking-wide mb-4">Previous Question</h2>
          
          <div class="space-y-4">
            <div class="p-4 rounded-lg" :class="[
              history[0].correct 
                ? 'bg-green-500/10 border-2 border-green-500/20'
                : 'bg-red-500/10 border-2 border-red-500/20'
            ]">
              <p class="text-xl mb-3 text-gray-200">{{ history[0].question }}</p>
              <p class="font-medium text-2xl mb-3" :class="[
                history[0].correct 
                  ? 'text-green-400'
                  : 'text-red-400'
              ]">
                {{ history[0].answer }}
                <span class="ml-2">{{ history[0].correct ? '✓' : '✗' }}</span>
              </p>
              <p class="font-medium text-lg" :class="[
                history[0].correct 
                  ? 'text-green-400'
                  : 'text-red-400'
              ]">
                {{ history[0].feedback }}
              </p>
            </div>
            
            <div class="mt-4 p-4 bg-blue-500/10 border-2 border-blue-500/20 rounded-lg">
              <h3 class="text-sm uppercase tracking-wide text-blue-400 mb-2">Did you know?</h3>
              <p class="text-gray-300 text-lg">{{ history[0].tidbit }}</p>
            </div>
          </div>
        </section>
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

// Focus input whenever it's available or when loading completes
watch([currentQuestion, loading], async ([newQuestion, isLoading]) => {
  if (newQuestion && !isLoading) {
    await nextTick()
    answerInput.value?.focus()
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

/* Gradient text animation */
@keyframes gradient {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

h1.bg-gradient-to-r {
  background-size: 200% auto;
  animation: gradient 8s ease infinite;
}

/* Input focus animation */
input:focus {
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
  transition: all 0.2s ease-in-out;
}

/* Button transitions */
button {
  transition: opacity 0.2s ease;
}

button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Add these new styles */
section {
  transition: all 0.3s ease;
}

.border-2 {
  transition: border-color 0.3s ease;
}
</style>
