<template>
  <div
    class="happyfishing"
    :class="{
      'happyfishing-fish-1': props.name === '1',
      'happyfishing-fish-2': props.name === '2',
      'happyfishing-fish-3': props.name === '3',
      'happyfishing-fish-4': props.name === '4',
      'happyfishing-fish-5': props.name === '5',
      'happyfishing-fish-6': props.name === '6',
      'happyfishing-fish-7': props.name === '7',
      'happyfishing-fish-8': props.name === '8',
      'happyfishing-fish-9': props.name === '9',
      'happyfishing-fish-10': props.name === '10',
      'happyfishing-fish-11': props.name === '11',
      'happyfishing-fish-12': props.name === '12',
      'happyfishing-fish-13': props.name === '13',
      'happyfishing-fish-14': props.name === '14',
      'happyfishing-fish-15': props.name === '15',
      'happyfishing-fish-16': props.name === '16',
      'happyfishing-fish-17': props.name === '17',
      'happyfishing-fish-18': props.name === '18',
      'happyfishing-fish-19': props.name === '19',
      'happyfishing-fish-20': props.name === '20',
      'happyfishing-fish-21': props.name === '21',
      'happyfishing-fish-22': props.name === '22',
      'happyfishing-fish-23': props.name === '23',
      'happyfishing-fish-24': props.name === '24',
      'happyfishing-fish-25': props.name === '25' || props.name == 'f1',
      'happyfishing-fish-26': props.name === '26' || props.name == 'f2',
      'happyfishing-fish-27': props.name === '27' || props.name == 'f3',
      'happyfishing-fish-28': props.name === '28',
      'happyfishing-fish-29': props.name === '29',
      'happyfishing-fish-30': props.name === '30',
      'happyfishing-fish-31': props.name === '31',
      'happyfishing-fish-32': props.name === '32',
      'happyfishing-fish-33': props.name === '33',
      'happyfishing-fish-34': props.name === '34',
      'happyfishing-fish-35': props.name === '35',
      'happyfishing-skill-s1': props.name === 's1',
      'happyfishing-skill-s2': props.name === 's2',
    }"
  ></div>
</template>

<script setup>
const props = defineProps({
  name: {
    type: String,
    default: '',
  },
})
</script>
