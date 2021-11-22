<template>
    <v-data-table 
    :headers="headers"
    :items="atencions"
    :search="search"
    sort-by="id_persona"
    class="elevation-1" style="width:1300px">
        <template v-slot:top>
            <v-toolbar flat color="white">
                <v-toolbar-title>DATASET - ATENCIÓN COVID LIST</v-toolbar-title>
                <v-divider class="mx-4" inset vertical />
                <v-spacer></v-spacer>
                <v-text-field label="Search Atención" append-icon="search" 
                class="text-xs-center" v-model="search" single-line hide-details></v-text-field>
                <v-spacer></v-spacer>
                <v-dialog v-model="dialog" max-width="500px">
                    <template v-slot:activator="{ on }">
                        <v-btn color="primary" dark class="mb-2" v-on="on">Nueva Atención</v-btn>
                    </template>
                    <v-card>
                        <v-card-title>
                            <span class="headline">{{ formTitle }}</span>
                        </v-card-title>

                        <v-card-text>
                            <v-container>
                                <v-row>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="id_persona" label="Id_persona"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="id_eess" label="Id_eess"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="fecha_ingreso" label="Fecha_ingreso"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="hora_ingreso" label="Hora_ingreso"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="es_recuperado" label="Es_recuperado"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="fecha_alta" label="Fecha_alta"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="es_recuperado_voluntario" label="Es_recuperado_voluntario"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="fecha_alta_voluntaria" label="Fecha_alta_voluntaria"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="es_fallecido" label="Es_fallecido"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="fecha_fallecido" label="Fecha_fallecido"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="es_referido" label="Es_referido"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="fecha_referido" label="Fecha_referido"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="12" md="12">
                                        <v-text-field v-model="eess_destino_id" label="Eess_destino_id"></v-text-field>
                                    </v-col>
                                </v-row>
                            </v-container>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
                            <v-btn color="blue darken-1" text @click="save">Save</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </v-toolbar>
        </template>
        <template v-slot:item="{ item }">
            <tr>
                <td>{{ item.id_persona }}</td>
                <td>{{ item.id_eess }}</td>
                <td>{{ item.fecha_ingreso }}</td>
                <td>{{ item.hora_ingreso }}</td>
                <td>{{ item.es_recuperado }}</td>
                <td>{{ item.fecha_alta }}</td>
                <td>{{ item.es_recuperado_voluntario }}</td>
                <td>{{ item.fecha_alta_voluntaria }}</td>
                <td>{{ item.es_fallecido }}</td>
                <td>{{ item.fecha_fallecido }}</td>
                <td>{{ item.es_referido }}</td>
                <td>{{ item.fecha_referido }}</td>
                <td>{{ item.eess_destino_id }}</td>
            </tr>
        </template>

        <template v-slot:no-data>
            <v-btn color="primary" @click="listAtencions">Reset</v-btn>
        </template>
    </v-data-table>
</template>

<script>
    import axios from 'axios'
    
    export default {
        data:() => ({
            search: '',
            id_persona: '',
            id_eess: '',
            fecha_ingreso: '',
            hora_ingreso: '',
            es_recuperado: '',
            fecha_alta: '',
            es_recuperado_voluntario: '',
            fecha_alta_voluntaria: '',
            es_fallecido: '',
            fecha_fallecido: '',
            es_referido: '',
            fecha_referido: '',
            eess_destino_id: '',
            dialog: false,
            atencions: [],
            valid: 0,
            validMessage: [],
            headers: [
                { text: 'Id_persona', value: 'Id_persona', sortable: true },
                { text: 'Id_eess', value: 'id_eess', sortable: false },
                { text: 'Fecha_ingreso', value: 'fecha_ingreso', sortable: false },
                { text: 'Hora_ingreso', value: 'hora_ingreso', sortable: false },
                { text: 'Es_recuperado', value: 'es_recuperado', sortable: false },
                { text: 'Fecha_alta', value: 'fecha_alta', sortable: false },
                { text: 'Es_recuperado_voluntario', value: 'es_recuperado_voluntario', sortable: false },
                { text: 'Fecha_alta_voluntaria', value: 'fecha_alta_voluntaria', sortable: false },
                { text: 'Es_fallecido', value: 'es_fallecido', sortable: false },
                { text: 'Fecha_fallecido', value: 'fecha_fallecido', sortable: false },
                { text: 'Es_referido', value: 'es_referido', sortable: false },
                { text: 'Fecha_referido', value: 'fecha_referido', sortable: false },
                { text: 'Eess_destino_id', value: 'eess_destino_id', sortable: false },
            ]
        }),
        computed: {
            formTitle() {
                return'New Atencion';
            }
        },
        watch: {
            dialog (val) {
                val || this.close()
            }
        },
        created() {
            this.listAtencions();
        },
        methods: {
            listAtencions() {
                let me= this;
                axios.get('atencions').then(function(response){
                    me.atencions = response.data;
                }).catch(function(error){
                    console.log(error);
                })
            },
            close() {
                this.dialog = false;
            },
            clean() {
                this.id_persona = "",
                this.id_eess = "",
                this.fecha_ingreso = "",
                this.hora_ingreso = "",
                this.es_recuperado = "",
                this.fecha_alta = "",
                this.es_recuperado_voluntario = "",
                this.fecha_alta_voluntaria = "",
                this.es_fallecido = "",
                this.fecha_fallecido = "",
                this.es_referido = "",
                this.fecha_referido = "",
                this.eess_destino_id = ""
            },
            save() {
                let me=this;
                axios.post('atencions', {
                    'id_persona' : me.id_persona,
                    'id_eess': me.id_eess,
                    'fecha_ingreso': me.fecha_ingreso,
                    'hora_ingreso': me.hora_ingreso,
                    'es_recuperado': me.es_recuperado,
                    'fecha_alta': me.fecha_alta,
                    'es_recuperado_voluntario': me.es_recuperado_voluntario,
                    'fecha_alta_voluntaria': me.fecha_alta_voluntaria,
                    'es_fallecido': me.es_fallecido,
                    'fecha_fallecido': me.fecha_fallecido,
                    'es_referido': me.es_referido,
                    'fecha_referido': me.fecha_referido,
                    'eess_destino_id': me.eess_destino_id,
                }).then(function(response){
                    me.close();
                    me.listAtencions();
                    me.clean();
                }).catch(function(error) {
                    console.log(error);
                });
            }
        }
    }
</script>