import fetch from 'node-fetch'
import assert from "assert";


describe('web test', function () {

    const url_admin_num = "http://localhost:7890/admin/number"
    let contract_name="DEYI_NFTAsset"
    // let url_set_contract_name = "http://localhost:7890/contract/name/" + contract_name
    const url_get_contract_name = "http://localhost:7890/contract/name"

    const test_username_admin1 = "admin1"
    const admin1_addr = "b2cdd0e4f79b1fcd000f7e671e16a2b5b6b7602d"
    const test_username_user1 = "user1"
    let user1_addr
    const test_username_user2 = "user2"
    let user2_addr
    let test_username_newuser = "user" + Math.floor(Math.random()*1000000000000+100).toString()
    const url_client = "http://localhost:7890/client/"
    const url_client_pk = "http://localhost:7890/user/pk"
    const url_client_addr = "http://localhost:7890/user/addr"
    const url_new_user = "http://localhost:7890/user/new/"

    const asset_name = "TEST_NAME"
    const asset_symbol = "TEST_SYMBOL"
    const asset_uri = "TEST_URI"
    let asset_build_version="24.0." + Math.floor(Math.random()*1000000000000+1000000000).toString()
    let asset_update_version="32.0." + Math.floor(Math.random()*1000000000000+1000000000).toString()

    const url_build_contract = "http://localhost:7890/contract/build"
    const url_update_contract = "http://localhost:7890/contract/update/"
    const url_freeze_contract = "http://localhost:7890/contract/freeze"
    const url_unfreeze_contract = "http://localhost:7890/contract/unfreeze"
    // const url_revoke_contract = "http://localhost:7890/contract/revoke"

    const url_get_asset_name = "http://localhost:7890/contract/asset/name"
    const url_get_asset_tokenuri = "http://localhost:7890/contract/asset/tokenuri"
    const url_get_asset_symbol = "http://localhost:7890/contract/asset/symbol"

    const url_balance = "http://localhost:7890/contract/asset/balance/"
    const url_owner = "http://localhost:7890/contract/asset/owner/"
    
    const url_mint = "http://localhost:7890/contract/asset/mint "
    const test_tokenid1 = Math.floor(Math.random()*1000000000000+1000000000).toString()
    const test_tokenid2 = Math.floor(Math.random()*1000000000000+1000000000).toString()
    const test_tokenid3 = Math.floor(Math.random()*1000000000000+1000000000).toString()
    

    const url_transfer = "http://localhost:7890/contract/asset/transfer"

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
        // let response = await fetch(url_set_contract_name, {
        //     method: "POST",
        //     headers: {
        //         "Content-type": "application/json; charset=UTF-8"
        //     }
        // })

        let response = await fetch(url_get_contract_name);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.name,contract_name)
        } else {
          alert("HTTP-Error: " + response.status);
        }
    });

    it('set client test', async () => {
        let response = await fetch(url_client+test_username_admin1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_client_addr);
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

        response = await fetch(url_client+test_username_user1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_client_addr);
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

        response = await fetch(url_client+test_username_user2, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_client_addr);
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

    it('build new a contract', async () => {

        let response = await fetch(url_client+test_username_admin1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_build_contract, {
            method: "POST",
            body: JSON.stringify({
                version: asset_build_version,
                name: asset_name,
                symbol: asset_symbol,
                tokenuri: asset_uri,
                opadmin: admin1_addr,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_get_asset_name);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.name,asset_name)
        } else {
          alert("HTTP-Error: " + response.status);
        }

        response = await fetch(url_get_asset_symbol);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.symbol,asset_symbol)
        } else {
          alert("HTTP-Error: " + response.status);
        }

        response = await fetch(url_get_asset_tokenuri);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.tokenuri,asset_uri+"/0")
        } else {
          alert("HTTP-Error: " + response.status);
        }        

    });

    it('tokens mint test', async () => {

        let response = await fetch(url_client+test_username_admin1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_mint, {
            method: "POST",
            body: JSON.stringify({
                address: user1_addr,
                tokenid: test_tokenid1,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_mint, {
            method: "POST",
            body: JSON.stringify({
                address: user1_addr,
                tokenid: test_tokenid2,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })
        
        response = await fetch(url_mint, {
            method: "POST",
            body: JSON.stringify({
                address: user1_addr,
                tokenid: test_tokenid3,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_balance+user1_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,3)
        } else {
          alert("HTTP-Error: " + response.status);
        }     

        response = await fetch(url_balance+user2_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,0)
        } else {
          alert("HTTP-Error: " + response.status);
        }  

        response = await fetch(url_owner+test_tokenid1);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_owner+test_tokenid2);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_owner+test_tokenid3);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

    });


    it('tokens transfer token1 test', async () => {

        let response = await fetch(url_client+test_username_user1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_transfer, {
            method: "POST",
            body: JSON.stringify({
                to: user2_addr,
                tokenid: test_tokenid1,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_balance+user1_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,2)
        } else {
          alert("HTTP-Error: " + response.status);
        }     

        response = await fetch(url_balance+user2_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,1)
        } else {
          alert("HTTP-Error: " + response.status);
        }  

        response = await fetch(url_owner+test_tokenid1);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user2_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_owner+test_tokenid2);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_owner+test_tokenid3);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user1_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

    });

    it('tokens transfer token2 and token3 test', async () => {

        let response = await fetch(url_client+test_username_user1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_transfer, {
            method: "POST",
            body: JSON.stringify({
                to: user2_addr,
                tokenid: test_tokenid2,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_transfer, {
            method: "POST",
            body: JSON.stringify({
                to: user2_addr,
                tokenid: test_tokenid3,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_balance+user1_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,0)
        } else {
          alert("HTTP-Error: " + response.status);
        }     

        response = await fetch(url_balance+user2_addr);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.balance,3)
        } else {
          alert("HTTP-Error: " + response.status);
        }  

        response = await fetch(url_owner+test_tokenid1);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user2_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_owner+test_tokenid2);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user2_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

        response = await fetch(url_owner+test_tokenid3);
        if (response.ok) { // if HTTP-status is 200-299
          // get the response body (the method explained below)
          let json = await response.json();
          assert.equal(json.owner, user2_addr)
        } else {
          alert("HTTP-Error: " + response.status);
        }             

    });

    it('contract management update/freeze/unfreeze/revoke', async () => {
        let response = await fetch(url_client+test_username_admin1, {
            method: "POST",
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })

        response = await fetch(url_update_contract+asset_update_version, {
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
        
        response = await fetch(url_freeze_contract, {
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
        
        response = await fetch(url_unfreeze_contract, {
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

        // response = await fetch(url_revoke_contract, {
        //     method: "POST",
        //     headers: {
        //         "Content-type": "application/json; charset=UTF-8"
        //     }
        // })
        // if (response.ok) { // if HTTP-status is 200-299
        //     // get the response body (the method explained below)
        //     let json = await response.json();
        //     console.log(json)
        // } else {
        //     alert("HTTP-Error: " + response.status);
        // }  

    });

});

