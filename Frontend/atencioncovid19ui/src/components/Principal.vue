<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/logo.svg')"
          class="my-3"
          contain
          height="150"
        />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          BD Atencion Covid 19 - TF Programacion Concurrente y Distribuida
        </h1>

        <h3 class="subheading font-weight-regular-bold mb-3">
          Sigi Andre Diaz Quiroz - U201621283 - Ingenieria de Software
        </h3>

        <h3 class="subheading font-weight-regular-bold mb-3">
          Nicole Jealid Kuylen Aranda - u201711586 - Ingenieria de Software
        </h3>

        <h3 class="subheading font-weight-regular-bold mb-3">
          Joaquin Alonso Egocheaga Ramos - U201715418 - Ingenieria de Software
        </h3>

        <h3 class="subheading font-weight-regular-bold mb-3">
          Jett Manuel Teixeira Walters – U201315252 - Ciencia de la Computación
        </h3>

        <p class="subheading font-weight-regular">
          Nuestro Backend consumiendo de nuestro DATASET.
          <br>¡Desde la comodidad de tu hogar!
          <a
            href="https://github.com/sigiandre/TF-Atencion-Covid-Concurrente"
            target="_blank"
          >Nuesto Backend</a>
        </p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'Principal',
    data: () => ({
    }),
  }
</script>