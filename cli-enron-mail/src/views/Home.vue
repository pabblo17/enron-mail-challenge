<template>

  <Navbar @searchTerm="eventSearch" />

  <div class="grid grid-cols-3 gap-4 py-8 px-8">
    <div class="col-span-2 bg-white rounded-xl shadow border">
      <div class="mt-5 px-5 py-5">
        <div class="bg-gray-500 p-5">
          <h4 class="text-xs font-medium text-white " v-if="isVisibleTotal">
            "<span class="inline-flex text-xs underline  font-bold leading-5">
              {{ termSearch }}
            </span>
            total found:
            <span class="inline-flex text-xs underline  font-bold leading-5">
              {{ total }}
            </span>
            "
          </h4>
        </div>
        <!-- TABLE OF RESULT -->
        <div class="flex flex-col mt-6">
          <div class="py-2 -my-2 overflow-x-auto sm:-mx-6 sm:px-6 lg:-mx-8 lg:px-8">
            <div class="inline-block min-w-full overflow-hidden border-b border-gray-200 shadow sm:rounded-lg">
              <table class="min-w-full">
                <thead>
                  <tr>
                    <th
                      class=" px-6 py-3 text-lt font-medium leading-4 tracking-wider text-left text-gray-500 uppercase bg-gray-200 border-b border-gray-200">
                      Subject
                    </th>
                    <th
                      class="px-6 py-3 text-lt font-medium leading-4 tracking-wider text-left text-gray-500 uppercase bg-gray-200 border-b border-gray-200">
                      From
                    </th>
                    <th
                      class="px-6 py-3 text-lt font-medium leading-4 tracking-wider text-left text-gray-500 uppercase bg-gray-200 border-b border-gray-200">
                      To
                    </th>
                    <th
                      class="px-6 py-3 text-lt font-medium leading-4 tracking-wider text-left text-gray-500 uppercase bg-gray-200 border-b border-gray-200">
                      Content
                    </th>
                    <th
                      class="px-6 py-3 text-lt font-medium leading-4 tracking-wider text-left text-gray-500 uppercase bg-gray-200 border-b border-gray-200">
                    </th>
                  </tr>
                </thead>

                <tbody class="bg-white align-text-top">
                  <tr v-for="item in result" :key="item">

                    <td class="px-4 py-1 border-b border-gray-500 ">
                      <div class="text-lt leading-5 text-gray-500 font-bold">
                        {{ item.subject }}
                      </div>
                    </td>

                    <td class="px-3 py-1 border-b border-gray-500 ">
                      <div class="text-lt leading-5 text-gray-500">
                        {{ item.from }}
                        
                      </div>
                    </td>

                    <td class="px-3 py-1 border-b border-gray-500 ">
                      <div class="overflow-auto h-32 w-52">
                        <p v-for="mailTo in item.to" :key="mailTo">
                          <span class="inline-flex text-lt0 leading-5 text-green-800 bg-green-100 rounded-full">
                            {{ mailTo }}
                          </span>
                        </p>
                      </div>
                    </td>
                    <td class="px-3 py-1 text-lt leading-5 text-gray-500 border-b border-gray-500 ">
                      <div v-html="item.highlight[0]"></div>
                    </td>
                    <td
                      class="px-4 py-1 text-lt font-medium leading-5 text-right border-b border-gray-500 whitespace-nowrap ">
                      <a @click="viewFull(item._id)" class="text-indigo-600 hover:text-indigo-900">View Full</a>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
      <ul class="flex mx-5 mb-6" v-if="isVisiblePaginator">
        <li class="mx-1 px-3 py-2 bg-gray-200 text-gray-700 hover:bg-gray-700 hover:text-gray-200 rounded-lg">
          <button class="flex items-center font-bold" @click="paginator(currentPage-1)" :disabled="currentPage-1 < 0" >previous</button>
        </li>
        <li class="mx-1 px-3 py-2 bg-gray-200 text-gray-700 hover:bg-gray-700 hover:text-gray-200 rounded-lg">
          <button class="font-bold" >{{currentPage+1}}</button>
        </li>
        <li class="mx-1 px-3 py-2 bg-gray-200 text-gray-700 hover:bg-gray-700 hover:text-gray-200 rounded-lg">
          <button class="flex items-center font-bold" @click="paginator(currentPage+1)" :disabled="currentPage == maxPage-1 ">Next</button>
        </li>
      </ul>
    </div>

    <div class="bg-white rounded-xl shadow border py-8 px-5">
      <div class="  borderborder-gray-500" v-if="isVisibleFullMail">
        <div class="bg-gray-500 p-5">
          <h4 class="text-xs font-bold leading-5 text-center text-white">Content Mail</h4>
        </div>
        <br>
        <div class="mx-4">
          <span class="inline-flex px-2 text-lt font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
            Date:
          </span>
          {{ mail.date }}
        </div>
        <div class="mx-4 ">
          <span class="inline-flex px-2 text-lt font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
            Subject:
          </span>
          {{ mail.subject }}
        </div>
        <div class="mx-4">
          <span class="inline-flex px-2 text-lt font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
            From:
          </span>
          {{ mail.from }}
        </div>
        <div class="mx-4">
          <span class="inline-flex px-2 text-lt font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
            To:
          </span>
          {{ mail.to.join('; ') }}
        </div>  <div class="mx-4">
          <span class="inline-flex px-2 text-lt font-semibold leading-5 text-green-800 bg-green-100 rounded-full">
            file:
          </span>
          {{ mail.path }}
        </div>

        <div class="mx-4 my-5">
          <p class="text-xs underline font-bold ">
            Message:
          </p>
          {{ mail.content }}
        </div>
      </div>
    </div>
  </div>

  <router-view />
</template>

<script>

import axios from 'axios';
import Navbar from '../components/Navbar';
import { ref } from "vue";

//import { useRoute } from "vue-router";

export default {
  components: {
    Navbar
  },
  setup() {
    const axiosInstance = axios.create({
      //baseURL: "http://localhost:8000/mail",
      headers: {
        "Content-type": "application/json",
        "index": process.env.VUE_APP_SEARCH_API_INDEX
      }
    })

    const getMailService = process.env.VUE_APP_SEARCH_API_SERVICE_FIND
    const searchMailService = process.env.VUE_APP_SEARCH_API_SERVICE_QUERY

    const isVisiblePaginator = ref(false)
    const isVisibleFullMail = ref(false)
    const isVisibleTotal = ref(false)

    const termSearch = ref("")
    const total = ref("0")
    const result = ref("")
    const mail = ref("")

    const currentPage = ref(0)
    const maxResult = ref(5)
    const maxPage = ref(0)


    const eventSearch = (term) => {
      termSearch.value = term
      axiosInstance.get(
        searchMailService, {
        params: {
          q: termSearch.value,
          page: "0",
          max_results: maxResult.value,
        }
      }
      )
        .then(response => {
          console.log(response.data)
          total.value = response.data.total
          if (total.value > 0) {
            result.value = response.data.mails
            isVisiblePaginator.value = true
            isVisibleTotal.value = true
            currentPage.value = 0
            var numberOfPages = total.value/maxResult.value
            maxPage.value = Math.trunc(numberOfPages)
            if ((Math.trunc(numberOfPages) - numberOfPages ) < 1){
              maxPage.value ++
            }
          } else {
            result.value = ""
            isVisiblePaginator.value = false
            isVisibleTotal.value = true
          }

        })
    }

    const viewFull = (itemId) => {
      axiosInstance.post(
        getMailService, {},
        {
          headers: { "id": itemId }
        }
      )
        .then(response => {
          console.log(response.data)
          mail.value = response.data
          isVisibleFullMail.value = true
        })
    }

    
    const paginator = (page) => {
      axiosInstance.get(
        searchMailService, {
        params: {
          q: termSearch.value,
          page: page.toString(),
          max_results: maxResult.value,
        }
      }
      )
        .then(response => {
          console.log(response.data)
          total.value = response.data.total
          if (total.value > 0) {
            result.value = response.data.mails
            currentPage.value = page
            isVisiblePaginator.value = true
            isVisibleTotal.value = true
          } else {
            result.value = ""
            isVisiblePaginator.value = false
            isVisibleTotal.value = true
          }
        })
    }

    return {
      termSearch,
      total,
      result,
      mail,
      currentPage,
      maxResult,
      maxPage,
      isVisibleTotal,
      isVisiblePaginator,
      isVisibleFullMail,
      eventSearch,
      viewFull,
      paginator,
    }

  }
}

</script>


<style lang="css">
mark::before,
mark::after {
  clip-path: inset(100%);
  clip: rect(1px, 1px, 1px, 1px);
  height: 1px;
  overflow: hidden;
  position: absolute;
  white-space: nowrap;
  width: 1px;
}

mark::before {
  content: " [highlight start] ";
}

mark::after {
  content: " [highlight end] ";
}
</style>