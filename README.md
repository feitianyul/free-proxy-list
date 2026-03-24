**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-24 12:41:12 UTC（2026-03-24 20:41:12 UTC+8）

**代理总数：206**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 205 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 206 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 8.217.106.71:8888 | ✓ 612ms | ✓ 1058ms | ✓ 675ms | ✓ 764ms | ✓ 617ms | http |
| 147.161.210.140:8800 | ✓ 1727ms | ✓ 1469ms | ✓ 791ms | ✓ 1042ms | ✓ 942ms | http |
| 167.103.115.102:8800 | ✓ 869ms | 否 | ✓ 1176ms | ✓ 1203ms | ✓ 1250ms | http |
| 167.103.34.108:8800 | ✓ 1594ms | 否 | ✓ 1342ms | ✓ 1519ms | ✓ 1626ms | http |
| 46.101.190.71:3128 | ✓ 1788ms | 否 | 否 | ✓ 1988ms | ✓ 1491ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1998ms | ✓ 1247ms | ✓ 1274ms | ✓ 1338ms | http |
| 155.212.132.241:3128 | ✓ 850ms | 否 | 否 | ✓ 1959ms | ✓ 1463ms | http |
| 160.238.65.9:3128 | 否 | ✓ 1443ms | ✓ 670ms | ✓ 1520ms | ✓ 1161ms | http |
| 160.238.65.3:3128 | 否 | 否 | ✓ 666ms | ✓ 1467ms | ✓ 1138ms | http |
| 160.238.65.2:3128 | 否 | 否 | ✓ 645ms | ✓ 1517ms | ✓ 1132ms | http |
| 160.238.65.6:3128 | 否 | 否 | ✓ 666ms | ✓ 1469ms | ✓ 1167ms | http |
| 160.238.65.8:3128 | 否 | 否 | ✓ 681ms | ✓ 1488ms | ✓ 1148ms | http |
| 160.238.65.5:3128 | 否 | 否 | ✓ 649ms | ✓ 1538ms | ✓ 1162ms | http |
| 167.103.31.122:8800 | ✓ 1331ms | 否 | ✓ 1260ms | 否 | ✓ 1548ms | http |
| 116.80.96.100:3172 | 否 | 否 | ✓ 1470ms | ✓ 1779ms | ✓ 1639ms | http |
| 120.92.212.16:8890 | ✓ 1134ms | ✓ 1380ms | ✓ 1128ms | ✓ 1162ms | ✓ 1148ms | http |
| 120.92.212.16:7890 | ✓ 1192ms | ✓ 1372ms | ✓ 1147ms | ✓ 1374ms | ✓ 920ms | http |
| 14.225.222.185:7890 | 否 | 否 | ✓ 1177ms | ✓ 1412ms | ✓ 1014ms | http |
| 147.161.239.240:8800 | ✓ 1149ms | 否 | ✓ 1270ms | ✓ 1788ms | 否 | http |
| 101.43.127.100:8877 | 否 | ✓ 1003ms | 否 | ✓ 1855ms | ✓ 1630ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 881ms | ✓ 1447ms | ✓ 1162ms | http |
| 222.228.171.92:8080 | ✓ 1914ms | 否 | 否 | ✓ 1136ms | ✓ 1015ms | http |
| 38.145.218.234:8451 | ✓ 910ms | 否 | ✓ 286ms | ✓ 745ms | ✓ 666ms | http |
| 35.225.22.61:80 | ✓ 926ms | 否 | ✓ 311ms | ✓ 1300ms | ✓ 988ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1646ms | ✓ 1202ms | ✓ 1360ms | http |
| 38.34.178.155:8448 | ✓ 771ms | ✓ 1044ms | ✓ 682ms | ✓ 959ms | ✓ 1186ms | http |
| 47.101.159.19:8899 | ✓ 897ms | ✓ 998ms | ✓ 897ms | ✓ 1086ms | ✓ 883ms | http |
| 8.222.175.80:6128 | ✓ 1275ms | 否 | ✓ 680ms | ✓ 1004ms | ✓ 830ms | http |
| 180.125.216.109:8118 | ✓ 1067ms | ✓ 1239ms | 否 | 否 | ✓ 846ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1998ms | ✓ 999ms | 否 | ✓ 1671ms | http |
| 218.89.134.230:3333 | ✓ 1529ms | ✓ 1679ms | 否 | 否 | ✓ 1949ms | http |
| 38.34.183.219:8446 | ✓ 665ms | 否 | ✓ 642ms | ✓ 682ms | ✓ 531ms | http |
| 38.34.179.173:8452 | ✓ 397ms | ✓ 1507ms | ✓ 305ms | ✓ 776ms | ✓ 542ms | http |
| 114.231.73.92:1080 | 否 | ✓ 1882ms | ✓ 1378ms | ✓ 1208ms | ✓ 911ms | http |
| 137.220.151.110:6005 | ✓ 1867ms | 否 | ✓ 926ms | ✓ 1047ms | ✓ 821ms | http |
| 210.223.44.230:3128 | ✓ 1560ms | ✓ 1272ms | ✓ 631ms | ✓ 1868ms | ✓ 689ms | http |
| 59.8.203.55:80 | ✓ 1250ms | ✓ 1396ms | ✓ 836ms | ✓ 1001ms | ✓ 738ms | http |
| 137.220.150.22:6005 | ✓ 1506ms | 否 | ✓ 1423ms | ✓ 1214ms | ✓ 1262ms | http |
| 150.241.71.15:1080 | 否 | 否 | ✓ 1517ms | ✓ 1621ms | ✓ 1768ms | http |
| 101.47.73.135:3128 | ✓ 1160ms | 否 | ✓ 1152ms | ✓ 1427ms | 否 | http |
| 38.34.179.46:8452 | ✓ 961ms | ✓ 832ms | ✓ 214ms | ✓ 1120ms | ✓ 1045ms | http |
| 38.34.179.78:8445 | ✓ 845ms | ✓ 1383ms | 否 | ✓ 905ms | ✓ 747ms | http |
| 45.136.130.191:8453 | ✓ 1703ms | 否 | ✓ 1205ms | 否 | ✓ 949ms | http |
| 38.34.179.60:8452 | ✓ 447ms | ✓ 1983ms | ✓ 123ms | ✓ 730ms | ✓ 611ms | http |
| 45.136.131.32:8445 | ✓ 420ms | ✓ 1228ms | ✓ 838ms | ✓ 893ms | ✓ 867ms | http |
| 38.145.218.227:8450 | ✓ 859ms | ✓ 1216ms | ✓ 191ms | ✓ 795ms | ✓ 579ms | http |
| 103.113.70.189:1081 | ✓ 1066ms | 否 | ✓ 772ms | ✓ 1186ms | ✓ 870ms | http |
| 38.34.179.174:8453 | ✓ 174ms | ✓ 755ms | ✓ 97ms | ✓ 730ms | ✓ 835ms | http |
| 38.34.179.81:8451 | ✓ 109ms | ✓ 818ms | ✓ 261ms | ✓ 732ms | ✓ 528ms | http |
| 38.34.179.86:8452 | ✓ 175ms | ✓ 799ms | ✓ 146ms | ✓ 806ms | ✓ 543ms | http |
| 38.34.179.88:8446 | ✓ 163ms | ✓ 1474ms | ✓ 147ms | ✓ 682ms | ✓ 534ms | http |
| 38.34.183.225:8450 | ✓ 282ms | ✓ 894ms | ✓ 290ms | ✓ 770ms | ✓ 506ms | http |
| 38.34.179.59:8444 | ✓ 232ms | ✓ 898ms | ✓ 828ms | ✓ 869ms | ✓ 663ms | http |
| 38.34.179.39:8452 | ✓ 283ms | 否 | ✓ 364ms | ✓ 715ms | ✓ 531ms | http |
| 38.34.183.222:8453 | ✓ 319ms | ✓ 1080ms | ✓ 83ms | ✓ 783ms | ✓ 532ms | http |
| 38.34.179.60:8450 | ✓ 242ms | ✓ 1386ms | ✓ 176ms | ✓ 686ms | ✓ 510ms | http |
| 38.34.179.204:8446 | ✓ 246ms | ✓ 1868ms | ✓ 699ms | ✓ 883ms | ✓ 636ms | http |
| 38.34.179.75:8453 | ✓ 180ms | ✓ 1947ms | ✓ 117ms | ✓ 725ms | ✓ 533ms | http |
| 144.31.79.117:8888 | ✓ 1289ms | 否 | ✓ 824ms | ✓ 1974ms | ✓ 1600ms | http |
| 38.34.179.165:8446 | ✓ 181ms | ✓ 1676ms | ✓ 422ms | ✓ 729ms | ✓ 504ms | http |
| 38.34.179.83:8448 | ✓ 1040ms | 否 | ✓ 436ms | ✓ 1363ms | ✓ 566ms | http |
| 38.145.220.8:8452 | ✓ 312ms | ✓ 1124ms | ✓ 599ms | ✓ 885ms | ✓ 1072ms | http |
| 194.67.99.223:1080 | ✓ 819ms | 否 | ✓ 1761ms | 否 | ✓ 1533ms | http |
| 38.34.183.221:8452 | ✓ 723ms | 否 | ✓ 682ms | ✓ 1575ms | ✓ 1304ms | http |
| 150.107.140.238:3128 | ✓ 1801ms | 否 | 否 | ✓ 1261ms | ✓ 973ms | http |
| 38.34.179.213:8450 | ✓ 259ms | ✓ 947ms | ✓ 1460ms | ✓ 1141ms | ✓ 578ms | http |
| 38.34.179.228:8453 | ✓ 456ms | 否 | ✓ 1161ms | ✓ 1552ms | 否 | http |
| 45.136.130.198:8449 | ✓ 224ms | 否 | ✓ 616ms | ✓ 667ms | ✓ 607ms | http |
| 45.136.131.35:8452 | ✓ 696ms | 否 | ✓ 80ms | ✓ 674ms | ✓ 909ms | http |
| 116.80.96.102:3172 | ✓ 1661ms | 否 | 否 | ✓ 1861ms | ✓ 1859ms | http |
| 62.113.119.14:8080 | ✓ 836ms | ✓ 1750ms | ✓ 1461ms | ✓ 1682ms | ✓ 1297ms | http |
| 38.34.179.203:8450 | ✓ 675ms | ✓ 1343ms | ✓ 657ms | ✓ 943ms | ✓ 582ms | http |
| 38.145.220.33:8448 | 否 | 否 | ✓ 90ms | ✓ 666ms | ✓ 496ms | http |
| 38.180.2.107:3128 | ✓ 1118ms | ✓ 1910ms | ✓ 1876ms | 否 | 否 | http |
| 38.34.179.97:8448 | ✓ 660ms | ✓ 641ms | ✓ 83ms | 否 | 否 | http |
| 38.34.179.61:8445 | ✓ 645ms | ✓ 706ms | ✓ 91ms | 否 | 否 | http |
| 38.34.179.100:8443 | ✓ 650ms | ✓ 791ms | ✓ 93ms | 否 | 否 | http |
| 38.145.208.242:8451 | ✓ 658ms | ✓ 788ms | ✓ 139ms | 否 | 否 | http |
| 38.34.179.88:8452 | ✓ 649ms | ✓ 807ms | ✓ 169ms | 否 | 否 | http |
| 38.145.218.230:8447 | ✓ 639ms | ✓ 794ms | ✓ 244ms | 否 | 否 | http |
| 45.136.131.62:8449 | ✓ 653ms | ✓ 973ms | ✓ 105ms | 否 | 否 | http |
| 38.34.179.34:8452 | ✓ 637ms | ✓ 814ms | ✓ 308ms | 否 | 否 | http |
| 45.136.131.68:8446 | ✓ 626ms | ✓ 1115ms | ✓ 163ms | 否 | 否 | http |
| 38.34.179.101:8446 | ✓ 679ms | ✓ 1188ms | ✓ 89ms | 否 | 否 | http |
| 45.136.130.178:8449 | ✓ 375ms | ✓ 1139ms | ✓ 96ms | ✓ 1075ms | ✓ 635ms | http |
| 45.136.131.66:8445 | ✓ 298ms | ✓ 682ms | ✓ 1419ms | ✓ 879ms | ✓ 554ms | http |
| 45.136.131.58:8449 | ✓ 305ms | ✓ 738ms | ✓ 1378ms | ✓ 1026ms | ✓ 602ms | http |
| 166.88.55.83:7890 | ✓ 731ms | ✓ 1064ms | ✓ 597ms | ✓ 760ms | ✓ 606ms | http |
| 38.145.218.228:8447 | ✓ 366ms | ✓ 1159ms | ✓ 590ms | ✓ 1144ms | ✓ 976ms | http |
| 1.231.81.166:3128 | ✓ 579ms | 否 | ✓ 547ms | ✓ 913ms | ✓ 744ms | http |
| 88.80.150.82:8080 | ✓ 1014ms | 否 | ✓ 960ms | ✓ 1816ms | ✓ 1460ms | https |
| 181.41.201.85:3128 | ✓ 783ms | 否 | ✓ 856ms | 否 | ✓ 1731ms | http |
| 38.145.218.229:8450 | ✓ 748ms | 否 | ✓ 1114ms | ✓ 1982ms | 否 | http |
| 38.145.220.198:8448 | 否 | ✓ 1179ms | ✓ 170ms | ✓ 1312ms | ✓ 1047ms | http |
| 45.136.130.173:8448 | 否 | 否 | ✓ 858ms | ✓ 1554ms | ✓ 1507ms | http |
| 38.145.220.9:8448 | 否 | ✓ 940ms | ✓ 397ms | ✓ 776ms | ✓ 532ms | http |
| 38.34.183.234:8450 | ✓ 1116ms | ✓ 871ms | ✓ 110ms | ✓ 682ms | ✓ 572ms | http |
| 38.145.203.135:8443 | ✓ 1052ms | ✓ 1590ms | 否 | ✓ 1101ms | ✓ 906ms | http |
| 38.34.179.82:8443 | ✓ 1699ms | ✓ 1983ms | ✓ 887ms | ✓ 1030ms | ✓ 861ms | http |
| 45.136.131.34:8448 | ✓ 326ms | ✓ 702ms | ✓ 171ms | ✓ 713ms | ✓ 543ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
