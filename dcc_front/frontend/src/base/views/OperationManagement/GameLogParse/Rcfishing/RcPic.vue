<template>
  <div
    class="rcfishing"
    :class="{
      'rcfishing-fish-1': props.name === '1',
      'rcfishing-fish-2': props.name === '2',
      'rcfishing-fish-3': props.name === '3',
      'rcfishing-fish-4': props.name === '4',
      'rcfishing-fish-5': props.name === '5',
      'rcfishing-fish-6': props.name === '6',
      'rcfishing-fish-7': props.name === '7',
      'rcfishing-fish-8': props.name === '8',
      'rcfishing-fish-9': props.name === '9',
      'rcfishing-fish-10': props.name === '10',
      'rcfishing-fish-11': props.name === '11',
      'rcfishing-fish-12': props.name === '12',
      'rcfishing-fish-13': props.name === '13',
      'rcfishing-fish-14': props.name === '14',
      'rcfishing-fish-15': props.name === '15',
      'rcfishing-fish-16': props.name === '16',
      'rcfishing-fish-17': props.name === '17',
      'rcfishing-fish-18': props.name === '18',
      'rcfishing-fish-19': props.name === '19',
      'rcfishing-fish-20': props.name === '20',
      'rcfishing-fish-21': props.name === '21',
      'rcfishing-fish-22': props.name === '22',
      'rcfishing-fish-23': props.name === '23',
      'rcfishing-fish-24': props.name === '24',
      'rcfishing-fish-25': props.name === '25',
      'rcfishing-fish-26': props.name === '26',
      'rcfishing-fish-27': props.name === '27' || props.name === 'ss',
      'rcfishing-fish-28': props.name === '28',
      'rcfishing-fish-29': props.name === '29',
      'rcfishing-fish-30': props.name === '30',
      'rcfishing-fish-31': props.name === '31',
      'rcfishing-skill-s1': props.name === 's1',
      'rcfishing-skill-s2': props.name === 's2',
      'rcfishing-skill-s3': props.name === 's3',
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
