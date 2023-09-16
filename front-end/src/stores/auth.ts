import { defineStore } from "pinia";
import api from "./api";

interface AuthState {
  _name: string;
  _accessToken: string;
  _refreshToken: string;
  _isAuthenticated: boolean;
}

const authState: AuthState = {
  _name: "",
  _accessToken: "",
  _refreshToken: "",
  _isAuthenticated: false
};

export const useAuthStore = defineStore({
  id: "auth",
  state: (): AuthState => Object.assign(authState, JSON.parse(localStorage.getItem("authState")!)),

  actions: {
    async login(email: string, password: string): Promise<boolean> {
      const data = JSON.stringify({
        email: email,
        password: password
      });

      const config = {
        method: "post",
        url: "user/sign-in",
        headers: {
          "Content-Type": "application/json"
        },
        data: data
      };

      let retVal = false;

      await api
        .request(config)
        .then((response) => {
          if (response.status !== 200) {
            throw new Error("Error: " + response.status);
          }

          const data = response.data;
          this._name = data.user.name;
          this._isAuthenticated = true;
          this._accessToken = data.access_token;
          this._refreshToken = data.refresh_token;
          retVal = true;
          localStorage.setItem("authState", JSON.stringify(this.$state));
        })
        .catch((error) => {
          console.log(error);
        });

      return retVal;
    },

    async refresh() {
      const data = JSON.stringify({
        refresh_token: this._refreshToken
      });

      const config = {
        method: "post",
        maxBodyLength: Infinity,
        url: "/user/refresh",
        headers: {
          "Content-Type": "application/json"
        },
        data: data
      };

      await api
        .request(config)
        .then((response) => {
          if (response.status !== 200) {
            throw new Error("Error: " + response.status);
          }
          this._accessToken = response.data.access_token;
          localStorage.setItem("authState", JSON.stringify(this.$state));
        })
        .catch((error) => {
          console.log(error);
          this.logout();
        });
    },

    logout() {
      this._name = "";
      this._accessToken = "";
      this._refreshToken = "";
      this._isAuthenticated = false;
      localStorage.removeItem("authState");
    },

    afterUserCreate(params: any) {
      this._isAuthenticated = true;
      this._accessToken = params.access_token;
      this._refreshToken = params.refresh_token;
      this._name = params.user.name;
    }
  },

  getters: {
    user: (state) => state._name,
    accessToken: (state) => `Bearer ${state._accessToken}`,
    refreshToken: (state) => state._refreshToken,
    isAuthenticated: (state) => state._isAuthenticated
  }
});
