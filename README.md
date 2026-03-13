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

最后更新：2026-03-13 20:38:26 UTC（2026-03-14 04:38:26 UTC+8）

**代理总数：158**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 158 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 158 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.175:8443 | ✓ 761ms | ✓ 832ms | ✓ 791ms | ✓ 779ms | ✓ 517ms | http |
| 205.209.118.30:3138 | ✓ 1474ms | 否 | ✓ 1005ms | ✓ 1408ms | ✓ 1144ms | http |
| 113.160.132.26:8080 | ✓ 1786ms | ✓ 1311ms | 否 | ✓ 1553ms | ✓ 1092ms | http |
| 91.247.126.241:2080 | ✓ 1817ms | 否 | ✓ 1942ms | 否 | ✓ 1905ms | http |
| 103.87.67.75:3129 | ✓ 1631ms | 否 | ✓ 1669ms | 否 | ✓ 1216ms | http |
| 45.167.124.52:8080 | ✓ 1650ms | ✓ 1948ms | ✓ 1544ms | 否 | ✓ 1533ms | http |
| 120.232.242.119:22222 | ✓ 976ms | ✓ 1136ms | ✓ 893ms | ✓ 1081ms | ✓ 837ms | http |
| 186.148.180.46:999 | ✓ 1501ms | ✓ 1876ms | ✓ 1363ms | ✓ 1913ms | 否 | http |
| 45.168.238.193:8443 | 否 | ✓ 1103ms | ✓ 358ms | ✓ 1135ms | ✓ 930ms | http |
| 165.227.5.10:8888 | ✓ 1008ms | 否 | 否 | ✓ 1430ms | ✓ 632ms | http |
| 138.124.53.25:7443 | ✓ 755ms | 否 | 否 | ✓ 1776ms | ✓ 1587ms | http |
| 8.222.175.80:6128 | ✓ 1290ms | ✓ 1637ms | ✓ 684ms | ✓ 1069ms | ✓ 799ms | http |
| 120.240.35.173:22222 | ✓ 989ms | ✓ 1762ms | ✓ 1092ms | ✓ 1292ms | ✓ 1007ms | http |
| 35.225.22.61:80 | ✓ 1094ms | 否 | ✓ 1524ms | 否 | ✓ 1858ms | http |
| 120.240.29.51:22222 | ✓ 932ms | ✓ 1175ms | 否 | ✓ 1256ms | 否 | http |
| 45.136.131.63:8443 | ✓ 603ms | ✓ 639ms | ✓ 821ms | ✓ 682ms | ✓ 500ms | http |
| 45.136.131.47:8443 | ✓ 604ms | ✓ 637ms | ✓ 823ms | ✓ 704ms | ✓ 669ms | http |
| 81.70.169.194:80 | ✓ 870ms | ✓ 1192ms | ✓ 979ms | ✓ 1194ms | ✓ 916ms | http |
| 101.43.255.96:80 | ✓ 1054ms | ✓ 1164ms | ✓ 964ms | ✓ 1164ms | ✓ 1053ms | http |
| 162.240.154.26:3128 | ✓ 1872ms | 否 | ✓ 1332ms | ✓ 1356ms | 否 | http |
| 178.236.245.59:3128 | ✓ 1485ms | 否 | ✓ 1998ms | ✓ 1983ms | 否 | http |
| 120.238.159.189:22222 | ✓ 905ms | ✓ 1163ms | ✓ 960ms | ✓ 1099ms | ✓ 878ms | http |
| 59.46.216.131:30001 | ✓ 983ms | ✓ 1412ms | 否 | 否 | ✓ 957ms | http |
| 171.251.173.39:5104 | ✓ 1309ms | 否 | ✓ 1538ms | ✓ 1991ms | ✓ 1431ms | http |
| 137.184.6.117:3128 | 否 | 否 | ✓ 1188ms | ✓ 649ms | ✓ 652ms | http |
| 117.159.239.54:22222 | ✓ 852ms | 否 | ✓ 806ms | 否 | ✓ 792ms | http |
| 183.249.5.109:22222 | ✓ 842ms | 否 | ✓ 785ms | ✓ 954ms | ✓ 777ms | http |
| 222.184.48.235:22222 | 否 | ✓ 1122ms | ✓ 901ms | ✓ 1168ms | 否 | http |
| 120.55.163.237:10086 | ✓ 782ms | ✓ 996ms | ✓ 913ms | ✓ 1023ms | ✓ 907ms | http |
| 24.144.86.173:1080 | ✓ 1551ms | ✓ 1995ms | ✓ 1646ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1156ms | ✓ 1818ms | 否 | 否 | ✓ 977ms | http |
| 120.240.35.177:22222 | ✓ 1085ms | ✓ 1677ms | ✓ 1577ms | ✓ 1335ms | ✓ 1000ms | http |
| 103.86.131.62:80 | ✓ 964ms | 否 | 否 | ✓ 1305ms | ✓ 1231ms | http |
| 47.105.98.23:3128 | ✓ 1275ms | 否 | ✓ 1124ms | 否 | ✓ 1161ms | http |
| 24.199.124.152:3128 | ✓ 961ms | ✓ 774ms | ✓ 821ms | ✓ 677ms | ✓ 477ms | http |
| 120.92.212.16:8890 | ✓ 1818ms | 否 | ✓ 958ms | 否 | ✓ 1606ms | http |
| 120.238.159.230:22222 | ✓ 1560ms | 否 | ✓ 1160ms | ✓ 1655ms | 否 | http |
| 180.103.19.40:1080 | ✓ 1014ms | ✓ 1447ms | ✓ 1071ms | ✓ 1330ms | ✓ 1015ms | http |
| 129.213.162.27:17777 | ✓ 720ms | ✓ 1606ms | 否 | 否 | ✓ 1463ms | http |
| 120.198.141.80:22222 | ✓ 1009ms | ✓ 1223ms | ✓ 978ms | ✓ 1090ms | ✓ 909ms | http |
| 121.230.8.177:1080 | ✓ 1255ms | ✓ 1858ms | ✓ 1224ms | ✓ 1707ms | ✓ 1296ms | http |
| 3.79.194.222:32177 | ✓ 1730ms | 否 | ✓ 1590ms | 否 | ✓ 1410ms | http |
| 14.225.212.37:7890 | 否 | ✓ 1336ms | ✓ 1890ms | 否 | ✓ 842ms | http |
| 159.223.42.219:3128 | ✓ 712ms | 否 | ✓ 1406ms | ✓ 998ms | 否 | http |
| 8.219.97.248:80 | ✓ 1029ms | 否 | ✓ 1004ms | ✓ 1125ms | 否 | http |
| 62.60.177.204:34094 | ✓ 563ms | 否 | ✓ 1426ms | ✓ 1216ms | ✓ 1005ms | http |
| 120.198.141.84:22222 | 否 | ✓ 1244ms | ✓ 907ms | ✓ 1117ms | 否 | http |
| 113.59.32.162:22222 | ✓ 1393ms | ✓ 1496ms | ✓ 1067ms | ✓ 1188ms | ✓ 992ms | http |
| 106.117.208.101:7890 | ✓ 868ms | ✓ 1243ms | ✓ 988ms | ✓ 1239ms | ✓ 908ms | http |
| 45.186.6.104:3128 | ✓ 1954ms | ✓ 1799ms | ✓ 1857ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 925ms | ✓ 1732ms | ✓ 1717ms | 否 | 否 | http |
| 116.80.96.111:3172 | 否 | 否 | ✓ 1438ms | ✓ 1998ms | ✓ 1610ms | http |
| 111.79.111.126:3128 | ✓ 1312ms | 否 | ✓ 1313ms | ✓ 1863ms | 否 | http |
| 216.180.127.45:1080 | ✓ 731ms | 否 | ✓ 1359ms | ✓ 1467ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1270ms | 否 | ✓ 1372ms | 否 | ✓ 997ms | http |
| 194.5.212.40:8080 | ✓ 1663ms | ✓ 1963ms | ✓ 1199ms | 否 | ✓ 1525ms | http |
| 128.199.114.189:9090 | ✓ 1155ms | 否 | ✓ 1055ms | ✓ 1438ms | ✓ 903ms | http |
| 45.129.141.143:3128 | ✓ 1366ms | 否 | ✓ 1478ms | 否 | ✓ 1734ms | http |
| 103.39.51.190:8080 | ✓ 1742ms | 否 | 否 | ✓ 1615ms | ✓ 1448ms | http |
| 150.230.249.50:1080 | 否 | 否 | ✓ 1133ms | ✓ 867ms | ✓ 1689ms | http |
| 45.136.130.223:8443 | ✓ 781ms | ✓ 648ms | ✓ 85ms | ✓ 697ms | ✓ 540ms | http |
| 103.84.95.54:7890 | ✓ 624ms | 否 | ✓ 607ms | ✓ 1295ms | 否 | http |
| 120.240.35.175:22222 | ✓ 1053ms | ✓ 1476ms | ✓ 1302ms | ✓ 1251ms | ✓ 1042ms | http |
| 120.240.35.161:22222 | ✓ 1245ms | ✓ 1508ms | ✓ 992ms | ✓ 1294ms | ✓ 1117ms | http |
| 113.59.32.148:22222 | 否 | 否 | ✓ 1260ms | ✓ 1196ms | ✓ 1004ms | http |
| 222.184.48.248:22222 | ✓ 886ms | ✓ 1328ms | ✓ 798ms | ✓ 1187ms | ✓ 954ms | http |
| 121.126.185.63:25152 | ✓ 1847ms | 否 | ✓ 1438ms | 否 | ✓ 1402ms | http |
| 117.159.239.58:22222 | ✓ 937ms | ✓ 1003ms | ✓ 824ms | 否 | ✓ 791ms | http |
| 117.159.239.46:22222 | ✓ 842ms | ✓ 993ms | ✓ 773ms | ✓ 1033ms | ✓ 855ms | http |
| 117.159.239.49:22222 | ✓ 839ms | ✓ 991ms | ✓ 843ms | ✓ 1035ms | ✓ 832ms | http |
| 143.189.3.198:8080 | 否 | 否 | ✓ 792ms | ✓ 813ms | ✓ 1681ms | http |
| 120.238.159.228:22222 | 否 | ✓ 1197ms | ✓ 923ms | ✓ 1078ms | ✓ 841ms | http |
| 62.113.119.14:8080 | ✓ 902ms | 否 | ✓ 1337ms | ✓ 1673ms | ✓ 1259ms | http |
| 178.236.245.17:3128 | ✓ 773ms | 否 | ✓ 788ms | ✓ 1918ms | 否 | http |
| 38.145.220.39:8443 | ✓ 342ms | ✓ 596ms | ✓ 238ms | ✓ 665ms | ✓ 491ms | http |
| 38.145.220.81:8443 | ✓ 341ms | ✓ 590ms | ✓ 244ms | ✓ 669ms | ✓ 498ms | http |
| 38.145.218.206:8443 | ✓ 340ms | ✓ 635ms | ✓ 200ms | ✓ 666ms | ✓ 501ms | http |
| 38.145.220.8:8443 | ✓ 342ms | ✓ 637ms | ✓ 196ms | ✓ 669ms | ✓ 510ms | http |
| 38.145.220.20:8443 | ✓ 342ms | ✓ 605ms | ✓ 230ms | ✓ 666ms | ✓ 521ms | http |
| 38.34.183.16:8443 | ✓ 342ms | ✓ 637ms | ✓ 197ms | ✓ 687ms | ✓ 518ms | http |
| 38.145.220.82:8443 | ✓ 340ms | ✓ 594ms | ✓ 240ms | ✓ 697ms | ✓ 515ms | http |
| 38.145.220.79:8443 | ✓ 338ms | ✓ 642ms | ✓ 195ms | ✓ 693ms | ✓ 520ms | http |
| 38.145.220.196:8443 | ✓ 341ms | ✓ 801ms | ✓ 193ms | ✓ 691ms | ✓ 521ms | http |
| 38.145.220.56:8443 | ✓ 342ms | ✓ 793ms | ✓ 202ms | ✓ 705ms | ✓ 507ms | http |
| 38.145.203.35:8443 | ✓ 340ms | ✓ 1008ms | ✓ 141ms | ✓ 819ms | ✓ 492ms | http |
| 38.145.220.65:8443 | ✓ 337ms | ✓ 1153ms | ✓ 159ms | ✓ 669ms | ✓ 495ms | http |
| 38.145.203.39:8443 | ✓ 343ms | ✓ 1047ms | ✓ 101ms | ✓ 816ms | ✓ 506ms | http |
| 38.145.203.43:8443 | ✓ 341ms | ✓ 1007ms | ✓ 144ms | ✓ 814ms | ✓ 497ms | http |
| 38.145.218.113:8443 | ✓ 340ms | ✓ 1046ms | ✓ 103ms | ✓ 818ms | ✓ 516ms | http |
| 38.145.203.32:8443 | ✓ 343ms | ✓ 1020ms | ✓ 129ms | ✓ 820ms | ✓ 501ms | http |
| 38.145.203.19:8443 | ✓ 340ms | ✓ 1013ms | ✓ 141ms | ✓ 825ms | ✓ 515ms | http |
| 38.145.203.124:8443 | ✓ 355ms | ✓ 783ms | ✓ 200ms | ✓ 697ms | ✓ 803ms | http |
| 38.145.203.107:8443 | ✓ 356ms | ✓ 775ms | ✓ 205ms | ✓ 684ms | ✓ 820ms | http |
| 38.145.218.216:8443 | ✓ 356ms | ✓ 778ms | ✓ 203ms | ✓ 699ms | ✓ 815ms | http |
| 45.136.130.186:8443 | ✓ 341ms | ✓ 747ms | ✓ 560ms | ✓ 689ms | ✓ 512ms | http |
| 45.136.130.178:8443 | ✓ 346ms | ✓ 744ms | ✓ 561ms | ✓ 693ms | ✓ 507ms | http |
| 38.145.203.106:8443 | ✓ 354ms | ✓ 773ms | ✓ 364ms | ✓ 683ms | ✓ 676ms | http |
| 45.136.130.182:8443 | ✓ 340ms | ✓ 746ms | ✓ 562ms | ✓ 676ms | ✓ 525ms | http |
| 45.136.130.181:8443 | ✓ 339ms | ✓ 748ms | ✓ 563ms | ✓ 689ms | ✓ 509ms | http |
| 38.145.220.32:8443 | ✓ 358ms | ✓ 635ms | ✓ 711ms | ✓ 672ms | ✓ 492ms | http |

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
