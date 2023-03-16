import fetch from 'node-fetch'
import assert from "assert";



function delay(ms) {
  return new Promise((resolve) => {
    setTimeout(resolve, ms);
  });
}

describe('web test', function () {

    const common_base_uri = "http://localhost:7890"
    const nft_asset_base_uri = common_base_uri+"/nft"
    const multi_asset_base_uri = common_base_uri+"/multi"

    const url_admin_num = common_base_uri+ "/admin/number"

    const nft_contract_name="DDE_NFTAsset"
    const url_get_nft_contract_name = nft_asset_base_uri + "/contract/name"
    const multi_contract_name="DDE_MULTIAsset_TEST11"
    const url_get_multi_contract_name = multi_asset_base_uri + "/contract/name"

    const url_set_client = common_base_uri + "/client/"
    // const url_current_client_pk = common_base_uri + "/user/pk"
    const url_current_client_addr = common_base_uri + "/user/addr"
    const url_new_user = common_base_uri + "/user/new/"
    const url_height_by_txid = common_base_uri + "/height/"

    const test_username_admin1 = "admin1"
    const admin1_addr = "b2cdd0e4f79b1fcd000f7e671e16a2b5b6b7602d"
    const test_username_user1 = "user1"
    let user1_addr
    const test_username_user2 = "user2"
    let user2_addr
    let test_username_newuser = "user" + Math.floor(Math.random()*1000000000000+100).toString()
  

    let nft_asset_name = "TEST_NAME" + Math.floor(Math.random()*10000).toString()
    let nft_asset_symbol = "TEST_SYMBOL" + Math.floor(Math.random()*10000).toString()
    let nft_asset_uri = "http://test_uri" + Math.floor(Math.random()*10000).toString()
    // let asset_build_version="24.0." + Math.floor(Math.random()*1000000000000+1000000000).toString()
    let asset_update_version="32.0." + Math.floor(Math.random()*1000000000000+1000000000).toString()

    const url_get_nft_asset_name = nft_asset_base_uri + "/contract/asset/name"
    const url_get_nft_asset_symbol = nft_asset_base_uri + "/contract/asset/symbol"
    const url_get_nft_asset_uri = nft_asset_base_uri + "/contract/asset/uri/"

    const url_set_nft_asset_name = nft_asset_base_uri + "/contract/asset/name/"
    const url_set_nft_asset_symbol = nft_asset_base_uri + "/contract/asset/symbol/"
    const url_set_nft_asset_uri = nft_asset_base_uri + "/contract/asset/uri"

    const url_nft_asset_admins = nft_asset_base_uri + "/contract/asset/admins"
    const admin_example1 = ["b2cdd0e4f79b1fcd000f7e671e16a2b5b6b7602d","7c70d597d79c9cf7b9d6677397ab4e866e31978f"];
    const admin_example2 = ["b2cdd0e4f79b1fcd000f7e671e16a2b5b6b7602d", "586b833a299cca32deda1d9f85ecbef02703a436", "7c70d597d79c9cf7b9d6677397ab4e866e31978f"];

    // const url_build_contract = "http://localhost:7890/erc721/contract/build"
    const url_nft_update_contract = nft_asset_base_uri + "/contract/update/"
    const url_nft_freeze_contract = nft_asset_base_uri + "/contract/freeze"
    const url_nft_unfreeze_contract = nft_asset_base_uri + "/contract/unfreeze"
    // const url_revoke_contract = "http://localhost:7890/contract/revoke"

    const url_nft_get_asset_name = nft_asset_base_uri + "/contract/asset/name"
    const url_nft_get_asset_tokenuri = nft_asset_base_uri + "/contract/asset/tokenuri"
    const url_nft_get_asset_symbol = nft_asset_base_uri + "/contract/asset/symbol"
    const url_nft_balance = nft_asset_base_uri+"/contract/asset/balance/"
    const url_nft_owner = nft_asset_base_uri+"/contract/asset/owner/"
    const url_nft_mint = nft_asset_base_uri + "/contract/asset/mint "
    const test_tokenid1 = Math.floor(Math.random()*1000000000000+1000000000).toString()
    const test_tokenid2 = Math.floor(Math.random()*1000000000000+1000000000).toString()
    const test_tokenid3 = Math.floor(Math.random()*1000000000000+1000000000).toString()

    const url_nft_async_mint = nft_asset_base_uri + "/contract/asset/async/mint "
    const test_tokenid4 = Math.floor(Math.random()*1000000000000+1000000000).toString()
    const test_tokenid5 = Math.floor(Math.random()*1000000000000+1000000000).toString()
    const test_tokenid6 = Math.floor(Math.random()*1000000000000+1000000000).toString()

    const url_nft_transfer = nft_asset_base_uri+"/contract/asset/transfer"
    const url_nft_async_transfer = nft_asset_base_uri+"/contract/asset/async/transfer"


    const url_get_multi_asset_uri = multi_asset_base_uri + "/contract/asset/uri/"
    const url_set_multi_asset_uri = multi_asset_base_uri + "/contract/asset/uri"
    const url_multi_asset_admins = multi_asset_base_uri + "/contract/asset/admins"

    const url_multi_balance = multi_asset_base_uri+"/contract/asset/balance?"
    const url_multi_batch_balance = multi_asset_base_uri+"/contract/asset/batchbalance"
    const url_multi_mint = multi_asset_base_uri + "/contract/asset/mint"
    const url_multi_async_mint = multi_asset_base_uri + "/contract/asset/async/mint"
    const url_multi_batchmint = multi_asset_base_uri + "/contract/asset/batchmint"
    const url_multi_async_batchmint = multi_asset_base_uri + "/contract/asset/async/batchmint"

    const url_multi_asset_transfer = multi_asset_base_uri + "/contract/asset/transfer"
    const url_multi_asset_batchtransfer = multi_asset_base_uri + "/contract/asset/batchtransfer"
    const url_multi_asset_async_transfer = multi_asset_base_uri + "/contract/asset/async/transfer"
    const url_multi_asset_async_batchtransfer = multi_asset_base_uri + "/contract/asset/async/batchtransfer"

    it('get admin group numbers response', async () => {
        let response = await fetch(url_admin_num);

        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.number,4)
        } else {
          alert("HTTP-Error: " + response.status);
        }
    });

    it('get contract names test', async () => {
        let response = await fetch(url_get_nft_contract_name);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.name,nft_contract_name)
        } else {
          alert("HTTP-Error: " + response.status);
        }

        response = await fetch(url_get_multi_contract_name);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.name,multi_contract_name)
        } else {
          alert("HTTP-Error: " + response.status);
        }

        
    });

    it('set client test', async () => {
        let response = await fetch(url_set_client + test_username_admin1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_current_client_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.address,admin1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }
    });

    it('gen new user test', async () => {

        let response = await fetch(url_new_user+test_username_user1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_set_client+test_username_user1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_current_client_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.notEqual(json.address,"")
          user1_addr=json.address
          console.log("user1 addr: ",user1_addr)
        }

        response = await fetch(url_new_user+test_username_user2, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_set_client+test_username_user2, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_current_client_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.notEqual(json.address,"")
          user2_addr=json.address
          console.log("user2 addr: ",user2_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }
  
    });

    it('nft name/symbol/url test', async () => {

      let response = await fetch(url_set_client+test_username_admin1, {
          method: "POST",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      response = await fetch(url_set_nft_asset_name + nft_asset_name, {
          method: "POST",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      response = await fetch(url_set_nft_asset_symbol + nft_asset_symbol, {
        method: "POST",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_set_nft_asset_uri, {
        method: "POST",
        body: JSON.stringify({
            link: nft_asset_uri,
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_get_nft_asset_name);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.name,nft_asset_name)
      } else {
        alert("HTTP-Error: " + response.status);
      }

      response = await fetch(url_get_nft_asset_symbol);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.symbol,nft_asset_symbol)
      } else {
        alert("HTTP-Error: " + response.status);
      }

      response = await fetch(url_get_nft_asset_uri + "1");
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.tokenuri,nft_asset_uri+"/1")
      } else {
        alert("HTTP-Error: " + response.status);
      }        

    });

    it('nft admins test', async () => {

      let response = await fetch(url_set_client+test_username_admin1, {
          method: "POST",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      response = await fetch(url_nft_asset_admins, {
        method: "POST",
        body: JSON.stringify({
          adminlist: admin_example1,
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_nft_asset_admins);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert((JSON.parse(json.admins)).every((element, index) => {
          return element === admin_example1[index];}))
      } else {
        alert("HTTP-Error: " + response.status);
      }

      response = await fetch(url_nft_asset_admins, {
        method: "POST",
        body: JSON.stringify({
          adminlist: admin_example2,
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_nft_asset_admins);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert((JSON.parse(json.admins)).every((element, index) => {
          return element === admin_example2[index];}))
      } else {
        alert("HTTP-Error: " + response.status);
      }

    });

    it('nft tokens mint test', async () => {

        let response = await fetch(url_set_client+test_username_admin1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_nft_mint, {
            method: "POST",
            body: JSON.stringify({
                to: user1_addr,
                id: test_tokenid1,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_nft_mint, {
            method: "POST",
            body: JSON.stringify({
                to: user1_addr,
                id: test_tokenid2,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })
        
        response = await fetch(url_nft_mint, {
            method: "POST",
            body: JSON.stringify({
                to: user1_addr,
                id: test_tokenid3,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_nft_balance+user1_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,3)
        } else {
          alert("HTTP-Error: " + response.status);
        }     

        response = await fetch(url_nft_balance+user2_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,0)
        } else {
          alert("HTTP-Error: " + response.status);
        }  

        response = await fetch(url_nft_owner+test_tokenid1);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_nft_owner+test_tokenid2);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_nft_owner+test_tokenid3);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

    });

    it('nft tokens async-mint test', async () => {

      let response = await fetch(url_set_client+test_username_admin1, {
          method: "POST",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      let response_mint1 = await fetch(url_nft_async_mint, {
          method: "POST",
          body: JSON.stringify({
              to: user1_addr,
              id: test_tokenid4,
          }),
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })


      let response_mint2 = await fetch(url_nft_async_mint, {
          method: "POST",
          body: JSON.stringify({
              to: user1_addr,
              id: test_tokenid5,
          }),
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      let response_mint3 = await fetch(url_nft_async_mint, {
          method: "POST",
          body: JSON.stringify({
              to: user1_addr,
              id: test_tokenid6,
          }),
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })


      let txid1
      if (response_mint1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint1.json();
        txid1 = json.txid
      } else {
        alert("HTTP-Error: " + response_mint1.status);
      }   
      console.log("txid1:",txid1)
      let i = 0;
      while (i < 1) {
        await delay(1000)
        response_mint1 = await fetch(url_height_by_txid + txid1);
        if (response_mint1.height != "0") break;
        i++;
      }
      let height
      response_mint1 = await fetch(url_height_by_txid + txid1);
      if (response_mint1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint1.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_mint1.status);
      }   
      console.log("height:",height)     
      
      
      let txid2
      if (response_mint2.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint2.json();
        txid2 = json.txid
      } else {
        alert("HTTP-Error: " + response_mint2.status);
      }   
      console.log("txid2:",txid2)
      i = 0;
      while (i < 1) {
        await delay(1000)
        response_mint2 = await fetch(url_height_by_txid + txid2);
        if (response_mint2.height != "0") break;
        i++;
      }
      response_mint2 = await fetch(url_height_by_txid + txid2);
      if (response_mint2.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint2.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_mint2.status);
      }   
      console.log("height:",height)

      let txid3
      if (response_mint3.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint3.json();
        txid3 = json.txid
      } else {
        alert("HTTP-Error: " + response_mint3.status);
      }   
      console.log("txid3:",txid3)
      i = 0;
      while (i < 1) {
        await delay(1000)
        response_mint3 = await fetch(url_height_by_txid + txid3);
        if (response_mint3.height != "0") break;
        i++;
      }
      response_mint3 = await fetch(url_height_by_txid + txid3);
      if (response_mint3.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint3.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_mint3.status);
      }   
      console.log("height:",height)

      response = await fetch(url_nft_balance+user1_addr);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,6)
      } else {
        alert("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_nft_balance+user2_addr);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,0)
      } else {
        alert("HTTP-Error: " + response.status);
      }  

      response = await fetch(url_nft_owner+test_tokenid4);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.owner, user1_addr)
      } else {
        alert("HTTP-Error: " + response.status);
      }             

      response = await fetch(url_nft_owner+test_tokenid5);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.owner, user1_addr)
      } else {
        alert("HTTP-Error: " + response.status);
      }             

      response = await fetch(url_nft_owner+test_tokenid6);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.owner, user1_addr)
      } else {
        alert("HTTP-Error: " + response.status);
      }             

    });

    it('nft tokens transfer token1 test', async () => {

        let response = await fetch(url_set_client+test_username_user1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_nft_transfer, {
            method: "POST",
            body: JSON.stringify({
                to: user2_addr,
                id: test_tokenid1,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_nft_balance+user1_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,5)
        } else {
          alert("HTTP-Error: " + response.status);
        }     

        response = await fetch(url_nft_balance+user2_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,1)
        } else {
          alert("HTTP-Error: " + response.status);
        }  

        response = await fetch(url_nft_owner+test_tokenid1);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user2_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_nft_owner+test_tokenid2);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_nft_owner+test_tokenid3);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

    });

    it('nft tokens transfer token2 and token3 test', async () => {

        let response = await fetch(url_set_client+test_username_user1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_nft_transfer, {
            method: "POST",
            body: JSON.stringify({
                to: user2_addr,
                id: test_tokenid2,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_nft_transfer, {
            method: "POST",
            body: JSON.stringify({
                to: user2_addr,
                id: test_tokenid3,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_nft_balance+user1_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,3)
        } else {
          alert("HTTP-Error: " + response.status);
        }     

        response = await fetch(url_nft_balance+user2_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,3)
        } else {
          alert("HTTP-Error: " + response.status);
        }  

        response = await fetch(url_nft_owner+test_tokenid1);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user2_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_nft_owner+test_tokenid2);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user2_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_nft_owner+test_tokenid3);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user2_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

    });

    it('nft tokens async-transfer token1 test', async () => {

      let response = await fetch(url_set_client+test_username_user1, {
          method: "POST",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      response = await fetch(url_nft_async_transfer, {
          method: "POST",
          body: JSON.stringify({
              to: user2_addr,
              id: test_tokenid4,
          }),
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      let txid1
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        txid1 = json.txid
      } else {
        alert("HTTP-Error: " + response.status);
      }   
      console.log("txid1:",txid1)
      let i = 0;
      while (i < 1) {
        await delay(1000)
        response = await fetch(url_height_by_txid + txid1);
        if (response.height != "0") break;
        i++;
      }
      let height
      response = await fetch(url_height_by_txid + txid1);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response.status);
      }   
      console.log("height:",height)   

      response = await fetch(url_nft_balance+user1_addr);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,2)
      } else {
        alert("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_nft_balance+user2_addr);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,4)
      } else {
        alert("HTTP-Error: " + response.status);
      }  

      response = await fetch(url_nft_owner+test_tokenid4);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.owner, user2_addr)
      } else {
        alert("HTTP-Error: " + response.status);
      }             

      response = await fetch(url_nft_owner+test_tokenid5);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.owner, user1_addr)
      } else {
        alert("HTTP-Error: " + response.status);
      }             

      response = await fetch(url_nft_owner+test_tokenid6);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.owner, user1_addr)
      } else {
        alert("HTTP-Error: " + response.status);
      }             

    });

    it('nft tokens async-transfer token2 and token3 test', async () => {

      let response = await fetch(url_set_client+test_username_user1, {
          method: "POST",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      let response_transfer1 = await fetch(url_nft_async_transfer, {
          method: "POST",
          body: JSON.stringify({
              to: user2_addr,
              id: test_tokenid5,
          }),
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      let response_transfer2 = await fetch(url_nft_async_transfer, {
          method: "POST",
          body: JSON.stringify({
              to: user2_addr,
              id: test_tokenid6,
          }),
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })


      let txid1
      if (response_transfer1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_transfer1.json();
        txid1 = json.txid
      } else {
        alert("HTTP-Error: " + response_transfer1.status);
      }   
      console.log("txid1:",txid1)
      let i = 0;
      while (i < 1) {
        await delay(1000)
        response_transfer1 = await fetch(url_height_by_txid + txid1);
        if (response_transfer1.height != "0") break;
        i++;
      }
      let height
      response_transfer1 = await fetch(url_height_by_txid + txid1);
      if (response_transfer1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_transfer1.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_transfer1.status);
      }   
      console.log("height:",height)   

      let txid2
      if (response_transfer2.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_transfer2.json();
        txid1 = json.txid
      } else {
        alert("HTTP-Error: " + response_transfer2.status);
      }   
      console.log("txid1:",txid1)
      i = 0;
      while (i < 1) {
        await delay(1000)
        response_transfer2 = await fetch(url_height_by_txid + txid1);
        if (response_transfer2.height != "0") break;
        i++;
      }
      response_transfer2 = await fetch(url_height_by_txid + txid1);
      if (response_transfer2.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_transfer2.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_transfer2.status);
      }   
      console.log("height:",height) 


      response = await fetch(url_nft_balance+user1_addr);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,0)
      } else {
        alert("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_nft_balance+user2_addr);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,6)
      } else {
        alert("HTTP-Error: " + response.status);
      }  

      response = await fetch(url_nft_owner+test_tokenid4);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.owner, user2_addr)
      } else {
        alert("HTTP-Error: " + response.status);
      }             

      response = await fetch(url_nft_owner+test_tokenid5);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.owner, user2_addr)
      } else {
        alert("HTTP-Error: " + response.status);
      }             

      response = await fetch(url_nft_owner+test_tokenid6);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.owner, user2_addr)
      } else {
        alert("HTTP-Error: " + response.status);
      }             

    });

    it('nft contract management update/freeze/unfreeze', async () => {
        let response = await fetch(url_set_client+test_username_admin1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_nft_update_contract+asset_update_version, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })
        if (response.ok) { // if HTTP-status is 200-299
            // get the response body (the method explained below)
            let json = await response.json();
            console.log(json)
        } else {
            alert("HTTP-Error: " + response.status);
        }     
        
        response = await fetch(url_nft_freeze_contract, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })
        if (response.ok) { // if HTTP-status is 200-299
            // get the response body (the method explained below)
            let json = await response.json();
            console.log(json)
        } else {
            alert("HTTP-Error: " + response.status);
        }    
        
        response = await fetch(url_nft_unfreeze_contract, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })
        if (response.ok) { // if HTTP-status is 200-299
            // get the response body (the method explained below)
            let json = await response.json();
            console.log(json)
        } else {
            alert("HTTP-Error: " + response.status);
        }   

    });

    it('multi asset url test', async () => {

      let response = await fetch(url_set_client+test_username_admin1, {
          method: "POST",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      response = await fetch(url_set_multi_asset_uri, {
        method: "POST",
        body: JSON.stringify({
            link: nft_asset_uri,
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_get_multi_asset_uri + "1");
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.tokenuri,nft_asset_uri+"/1")
      } else {
        alert("HTTP-Error: " + response.status);
      }        

    });

    it('multi asset admins test', async () => {

      let response = await fetch(url_set_client+test_username_admin1, {
          method: "POST",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      response = await fetch(url_multi_asset_admins, {
        method: "POST",
        body: JSON.stringify({
          adminlist: admin_example1,
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_asset_admins);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert((JSON.parse(json.admins)).every((element, index) => {
          return element === admin_example1[index];}))
      } else {
        alert("HTTP-Error: " + response.status);
      }

      response = await fetch(url_multi_asset_admins, {
        method: "POST",
        body: JSON.stringify({
          adminlist: admin_example2,
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_asset_admins);
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert((JSON.parse(json.admins)).every((element, index) => {
          return element === admin_example2[index];}))
      } else {
        alert("HTTP-Error: " + response.status);
      }

    });

    it('mutli assets tokens mint test', async () => {

      let response = await fetch(url_set_client+test_username_admin1, {
          method: "POST",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      response = await fetch(url_multi_mint, {
        method: "POST",
        body: JSON.stringify({
          to: user1_addr,
          id: test_tokenid1,
          amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_mint, {
        method: "POST",
        body: JSON.stringify({
            to: user1_addr,
            id: test_tokenid2,
            amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_mint, {
        method: "POST",
        body: JSON.stringify({
            to: user1_addr,
            id: test_tokenid3,
            amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid1, {
          method: "GET",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid2, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid3, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     
    });

    it('mutli assets tokens async-mint test', async () => {

      let response = await fetch(url_set_client+test_username_admin1, {
          method: "POST",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })

      let response_mint1 = await fetch(url_multi_async_mint, {
        method: "POST",
        body: JSON.stringify({
          to: user1_addr,
          id: test_tokenid4,
          amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let response_mint2 = await fetch(url_multi_async_mint, {
        method: "POST",
        body: JSON.stringify({
            to: user1_addr,
            id: test_tokenid5,
            amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let response_mint3 = await fetch(url_multi_async_mint, {
        method: "POST",
        body: JSON.stringify({
            to: user1_addr,
            id: test_tokenid6,
            amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let txid1
      if (response_mint1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint1.json();
        txid1 = json.txid
      } else {
        alert("HTTP-Error: " + response_mint1.status);
      }   
      console.log("txid1:",txid1)
      let i = 0;
      while (i < 1) {
        await delay(1000)
        response_mint1 = await fetch(url_height_by_txid + txid1);
        if (response_mint1.height != "0") break;
        i++;
      }
      let height
      response_mint1 = await fetch(url_height_by_txid + txid1);
      if (response_mint1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint1.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_mint1.status);
      }   
      console.log("height:",height)     

      let txid2
      if (response_mint2.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint2.json();
        txid1 = json.txid
      } else {
        alert("HTTP-Error: " + response_mint2.status);
      }   
      console.log("txid1:",txid1)
      i = 0;
      while (i < 1) {
        await delay(1000)
        response_mint2 = await fetch(url_height_by_txid + txid1);
        if (response_mint2.height != "0") break;
        i++;
      }
      response_mint2 = await fetch(url_height_by_txid + txid1);
      if (response_mint2.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint2.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_mint2.status);
      }   
      console.log("height:",height) 


      let txid3
      if (response_mint3.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint3.json();
        txid1 = json.txid
      } else {
        alert("HTTP-Error: " + response_mint3.status);
      }   
      console.log("txid1:",txid1)
      i = 0;
      while (i < 1) {
        await delay(1000)
        response_mint3 = await fetch(url_height_by_txid + txid1);
        if (response_mint3.height != "0") break;
        i++;
      }
      response_mint3 = await fetch(url_height_by_txid + txid1);
      if (response_mint3.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint3.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_mint3.status);
      }   
      console.log("height:",height) 


      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid1, {
          method: "GET",
          headers: {
              "Content-type": "application/json; charset=UTF-8"
          }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid2, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid3, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     
    });

    it('mutli assets tokens batchmint test', async () => { 

      let response = await fetch(url_set_client+test_username_admin1, {
        method: "POST",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_batchmint, {
        method: "POST",
        body: JSON.stringify({
          to: user1_addr,
          ids: [test_tokenid1,test_tokenid2,test_tokenid3],
          amounts: ["100","100","100"],
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid1, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid2, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid3, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

    })

    it('mutli assets tokens async-batchmint test', async () => { 

      let response = await fetch(url_set_client+test_username_admin1, {
        method: "POST",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let response_mint1 = await fetch(url_multi_async_batchmint, {
        method: "POST",
        body: JSON.stringify({
          to: user1_addr,
          ids: [test_tokenid4,test_tokenid5,test_tokenid6],
          amounts: ["100","100","100"],
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let txid1
      if (response_mint1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint1.json();
        txid1 = json.txid
      } else {
        alert("HTTP-Error: " + response_mint1.status);
      }   
      console.log("txid1:",txid1)
      let i = 0;
      while (i < 1) {
        await delay(1000)
        response_mint1 = await fetch(url_height_by_txid + txid1);
        if (response_mint1.height != "0") break;
        i++;
      }
      let height
      response_mint1 = await fetch(url_height_by_txid + txid1);
      if (response_mint1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint1.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_mint1.status);
      }   
      console.log("height:",height)  

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid4, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid5, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid6, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     
    })

    it('mutli assets tokens batch transfer test', async () => { 

      let response = await fetch(url_set_client+test_username_user1, {
        method: "POST",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_asset_batchtransfer, {
        method: "POST",
        body: JSON.stringify({
          from: user1_addr,
          to: user2_addr,
          ids: [test_tokenid1.toString(),test_tokenid2.toString(),test_tokenid3.toString()],
          amounts: ["100","100","100"],
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid1, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid2, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid3, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }      
      
      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid1, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }         

      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid2, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid3, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     
    })

    it('mutli assets tokens async-batchtransfer test', async () => { 

      let response = await fetch(url_set_client+test_username_user1, {
        method: "POST",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let response_mint1 = await fetch(url_multi_asset_async_batchtransfer, {
        method: "POST",
        body: JSON.stringify({
          from: user1_addr,
          to: user2_addr,
          ids: [test_tokenid4.toString(),test_tokenid5.toString(),test_tokenid6.toString()],
          amounts: ["100","100","100"],
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let txid1
      if (response_mint1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint1.json();
        txid1 = json.txid
      } else {
        alert("HTTP-Error: " + response_mint1.status);
      }   
      console.log("txid1:",txid1)
      let i = 0;
      while (i < 1) {
        await delay(1000)
        response_mint1 = await fetch(url_height_by_txid + txid1);
        if (response_mint1.height != "0") break;
        i++;
      }
      let height
      response_mint1 = await fetch(url_height_by_txid + txid1);
      if (response_mint1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_mint1.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_mint1.status);
      }   
      console.log("height:",height)  

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid4, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid5, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid6, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }      
      
      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid4, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }         

      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid5, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid6, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,100)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

    })

    it('mutli assets tokens transfer test', async () => { 

      let response = await fetch(url_set_client+test_username_user1, {
        method: "POST",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_asset_transfer, {
        method: "POST",
        body: JSON.stringify({
          from: user1_addr,
          to: user2_addr,
          id: test_tokenid1.toString(),
          amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_asset_transfer, {
        method: "POST",
        body: JSON.stringify({
          from: user1_addr,
          to: user2_addr,
          id: test_tokenid2.toString(),
          amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_asset_transfer, {
        method: "POST",
        body: JSON.stringify({
          from: user1_addr,
          to: user2_addr,
          id: test_tokenid3.toString(),
          amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid1, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,0)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid2, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,0)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid3, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,0)
      } else {
        console.error("HTTP-Error: " + response.status);
      }      
      
      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid1, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }         

      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid2, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid3, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

    })

    it('mutli assets tokens async-transfertest', async () => { 

      let response = await fetch(url_set_client+test_username_user1, {
        method: "POST",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let response_tx1 = await fetch(url_multi_asset_async_transfer, {
        method: "POST",
        body: JSON.stringify({
          from: user1_addr,
          to: user2_addr,
          id: test_tokenid4.toString(),
          amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let response_tx2 = await fetch(url_multi_asset_async_transfer, {
        method: "POST",
        body: JSON.stringify({
          from: user1_addr,
          to: user2_addr,
          id: test_tokenid5.toString(),
          amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let response_tx3 = await fetch(url_multi_asset_async_transfer, {
        method: "POST",
        body: JSON.stringify({
          from: user1_addr,
          to: user2_addr,
          id: test_tokenid6.toString(),
          amount: "100",
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })

      let txid1
      if (response_tx1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_tx1.json();
        txid1 = json.txid
      } else {
        alert("HTTP-Error: " + response_tx1.status);
      }   
      console.log("txid1:",txid1)
      let i = 0;
      while (i < 1) {
        await delay(1000)
        response_tx1 = await fetch(url_height_by_txid + txid1);
        if (response_tx1.height != "0") break;
        i++;
      }
      let height
      response_tx1 = await fetch(url_height_by_txid + txid1);
      if (response_tx1.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_tx1.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_tx1.status);
      }   
      console.log("height:",height) 


      let txid2
      if (response_tx2.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_tx2.json();
        txid2 = json.txid
      } else {
        alert("HTTP-Error: " + response_tx2.status);
      }   
      console.log("txid2:",txid2)
      i = 0;
      while (i < 1) {
        await delay(1000)
        response_tx2 = await fetch(url_height_by_txid + txid2);
        if (response_tx2.height != "0") break;
        i++;
      }
      response_tx2 = await fetch(url_height_by_txid + txid2);
      if (response_tx2.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_tx2.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_tx2.status);
      }   
      console.log("height:",height) 


      let txid3
      if (response_tx3.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_tx3.json();
        txid3 = json.txid
      } else {
        alert("HTTP-Error: " + response_tx3.status);
      }   
      console.log("txid2:",txid3)
      i = 0;
      while (i < 1) {
        await delay(1000)
        response_tx3 = await fetch(url_height_by_txid + txid3);
        if (response_tx3.height != "0") break;
        i++;
      }
      response_tx3 = await fetch(url_height_by_txid + txid3);
      if (response_tx3.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response_tx3.json();
        height = json.height
      } else {
        alert("HTTP-Error: " + response_tx3.status);
      }   
      console.log("height:",height) 

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid4, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,0)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid5, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,0)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user1_addr + "&" + "id=" + test_tokenid6, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,0)
      } else {
        console.error("HTTP-Error: " + response.status);
      }      
      
      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid4, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }         

      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid5, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

      response = await fetch(url_multi_balance + "owner=" + user2_addr + "&" + "id=" + test_tokenid6, {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
      })
      if (response.ok) { // if HTTP-status is 200-299
        // get the response body (the method explained below)
        let json = await response.json();
        assert.equal(json.balance,200)
      } else {
        console.error("HTTP-Error: " + response.status);
      }     

    })

});

