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

最后更新：2026-03-10 22:30:35 UTC（2026-03-11 06:30:35 UTC+8）

**代理总数：72**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.191:8443 | ✓ 439ms | ✓ 850ms | ✓ 667ms | ✓ 761ms | ✓ 553ms | http |
| 45.136.130.175:8443 | ✓ 439ms | ✓ 670ms | ✓ 798ms | ✓ 776ms | ✓ 627ms | http |
| 45.136.131.63:8443 | ✓ 439ms | ✓ 1312ms | ✓ 714ms | ✓ 895ms | ✓ 560ms | http |
| 45.136.131.47:8443 | ✓ 438ms | 否 | ✓ 583ms | ✓ 757ms | ✓ 729ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1986ms | ✓ 1862ms | ✓ 1505ms | http |
| 217.76.245.80:999 | ✓ 721ms | ✓ 1251ms | ✓ 1377ms | ✓ 1831ms | ✓ 1516ms | http |
| 45.136.130.188:8443 | ✓ 425ms | ✓ 716ms | ✓ 140ms | ✓ 723ms | ✓ 558ms | http |
| 115.231.181.40:8128 | ✓ 944ms | ✓ 1383ms | ✓ 1702ms | 否 | ✓ 949ms | http |
| 193.233.85.17:3128 | ✓ 733ms | 否 | ✓ 1333ms | 否 | ✓ 1816ms | http |
| 138.124.53.25:7443 | ✓ 571ms | ✓ 1904ms | ✓ 1116ms | 否 | 否 | http |
| 190.9.109.198:999 | ✓ 937ms | ✓ 1581ms | ✓ 1334ms | ✓ 1678ms | ✓ 1399ms | http |
| 35.225.22.61:80 | ✓ 786ms | 否 | ✓ 1009ms | ✓ 1010ms | ✓ 784ms | http |
| 202.155.12.161:443 | 否 | ✓ 1571ms | ✓ 1182ms | ✓ 1057ms | 否 | http |
| 45.136.130.223:8443 | ✓ 561ms | ✓ 716ms | ✓ 307ms | ✓ 726ms | ✓ 550ms | http |
| 39.104.201.40:7890 | ✓ 917ms | ✓ 1215ms | ✓ 1003ms | ✓ 1246ms | ✓ 959ms | http |
| 35.183.64.191:29715 | ✓ 1699ms | 否 | ✓ 1139ms | 否 | ✓ 1915ms | http |
| 81.70.169.194:80 | ✓ 1177ms | ✓ 1289ms | ✓ 1249ms | ✓ 1318ms | 否 | http |
| 101.43.255.96:80 | ✓ 1293ms | 否 | ✓ 1482ms | ✓ 1456ms | ✓ 1288ms | http |
| 120.92.212.16:8890 | ✓ 1930ms | ✓ 1239ms | ✓ 998ms | 否 | ✓ 1241ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1376ms | ✓ 1205ms | ✓ 1693ms | http |
| 45.186.6.104:3128 | ✓ 1310ms | ✓ 1936ms | ✓ 1881ms | 否 | 否 | http |
| 152.70.98.46:8888 | ✓ 1593ms | ✓ 1180ms | ✓ 547ms | ✓ 792ms | ✓ 635ms | http |
| 158.69.185.37:3129 | ✓ 399ms | 否 | ✓ 885ms | ✓ 1101ms | ✓ 835ms | http |
| 114.55.226.123:10086 | ✓ 1063ms | ✓ 1407ms | ✓ 1032ms | 否 | ✓ 1043ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1182ms | ✓ 1273ms | ✓ 1154ms | http |
| 178.236.245.59:3128 | ✓ 1327ms | 否 | ✓ 1646ms | 否 | ✓ 1619ms | http |
| 178.236.245.17:3128 | ✓ 1341ms | 否 | ✓ 1646ms | 否 | ✓ 1614ms | http |
| 47.74.226.8:5001 | 否 | ✓ 1838ms | ✓ 957ms | 否 | ✓ 1069ms | http |
| 120.198.141.80:22222 | ✓ 1125ms | ✓ 1413ms | ✓ 1104ms | ✓ 1251ms | ✓ 971ms | http |
| 95.3.9.78:8080 | ✓ 868ms | 否 | ✓ 1244ms | 否 | ✓ 1468ms | http |
| 121.138.61.193:8442 | ✓ 1407ms | ✓ 1722ms | ✓ 1164ms | ✓ 1203ms | ✓ 882ms | http |
| 121.230.9.241:1080 | ✓ 1043ms | ✓ 1350ms | 否 | 否 | ✓ 1112ms | http |
| 121.230.8.72:1080 | ✓ 1183ms | ✓ 1537ms | ✓ 1197ms | ✓ 1910ms | 否 | http |
| 101.32.244.83:8080 | ✓ 975ms | 否 | ✓ 951ms | ✓ 1366ms | ✓ 1057ms | http |
| 121.43.196.210:8222 | ✓ 906ms | ✓ 1072ms | ✓ 910ms | ✓ 1200ms | ✓ 876ms | http |
| 121.43.196.213:8222 | ✓ 975ms | ✓ 1094ms | ✓ 862ms | ✓ 1136ms | ✓ 932ms | http |
| 47.77.193.180:1080 | ✓ 770ms | ✓ 1809ms | ✓ 498ms | ✓ 860ms | ✓ 713ms | http |
| 116.80.82.224:3172 | ✓ 1511ms | 否 | 否 | ✓ 1924ms | ✓ 1824ms | http |
| 38.180.2.107:3128 | ✓ 955ms | 否 | ✓ 1776ms | 否 | ✓ 1928ms | http |
| 129.213.162.27:17777 | ✓ 553ms | ✓ 1531ms | 否 | ✓ 1834ms | 否 | http |
| 205.209.118.30:3138 | ✓ 231ms | ✓ 1883ms | ✓ 207ms | 否 | ✓ 1288ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1375ms | ✓ 1180ms | ✓ 1419ms | ✓ 1175ms | http |
| 103.82.23.118:5196 | ✓ 1950ms | 否 | 否 | ✓ 1664ms | ✓ 1658ms | http |
| 1.225.116.115:1080 | ✓ 1462ms | ✓ 1426ms | 否 | ✓ 1895ms | ✓ 976ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 1075ms | ✓ 1899ms | ✓ 1269ms | http |
| 45.136.198.40:3128 | ✓ 1100ms | ✓ 1932ms | ✓ 1658ms | ✓ 1982ms | ✓ 1784ms | http |
| 121.204.158.249:3128 | 否 | ✓ 1339ms | 否 | ✓ 1764ms | ✓ 1778ms | http |
| 213.220.3.183:3128 | ✓ 1430ms | ✓ 1807ms | ✓ 1374ms | 否 | 否 | http |
| 46.183.25.8:443 | ✓ 923ms | 否 | ✓ 483ms | ✓ 1119ms | 否 | http |
| 152.42.213.210:8080 | ✓ 806ms | 否 | ✓ 1664ms | ✓ 1222ms | ✓ 947ms | http |
| 34.101.184.164:3128 | ✓ 991ms | 否 | 否 | ✓ 1418ms | ✓ 1070ms | http |
| 116.80.82.227:3172 | ✓ 1521ms | 否 | ✓ 1573ms | 否 | ✓ 1773ms | http |
| 120.232.242.119:22222 | ✓ 966ms | ✓ 1200ms | ✓ 972ms | ✓ 1128ms | ✓ 888ms | http |
| 222.184.48.235:22222 | ✓ 896ms | ✓ 1063ms | ✓ 949ms | ✓ 1133ms | ✓ 956ms | http |
| 120.198.141.75:22222 | ✓ 1179ms | 否 | ✓ 1140ms | ✓ 1335ms | ✓ 1092ms | http |
| 222.184.48.252:22222 | ✓ 1041ms | 否 | ✓ 1676ms | ✓ 1675ms | ✓ 1014ms | http |
| 222.184.48.248:22222 | ✓ 908ms | ✓ 1610ms | ✓ 936ms | 否 | ✓ 1001ms | http |
| 103.39.51.190:8080 | ✓ 1773ms | 否 | ✓ 1734ms | ✓ 1470ms | ✓ 1723ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1303ms | 否 | ✓ 892ms | ✓ 569ms | http |
| 67.169.98.211:443 | ✓ 657ms | 否 | ✓ 467ms | ✓ 1537ms | 否 | http |
| 103.35.188.243:3128 | ✓ 480ms | ✓ 1220ms | 否 | ✓ 1346ms | ✓ 1131ms | http |
| 120.238.159.227:22222 | ✓ 950ms | ✓ 1273ms | ✓ 967ms | ✓ 1101ms | ✓ 919ms | http |
| 157.245.194.13:8888 | ✓ 774ms | 否 | ✓ 1364ms | ✓ 1041ms | ✓ 860ms | http |
| 61.52.131.172:8443 | ✓ 864ms | ✓ 1146ms | ✓ 1549ms | ✓ 1174ms | 否 | http |
| 123.57.0.163:8888 | 否 | ✓ 1836ms | 否 | ✓ 1444ms | ✓ 1721ms | http |
| 190.212.131.238:3128 | ✓ 1759ms | 否 | ✓ 1801ms | 否 | ✓ 1792ms | http |
| 194.213.18.200:443 | ✓ 1545ms | 否 | ✓ 1893ms | 否 | ✓ 1898ms | http |
| 37.139.33.145:1080 | ✓ 1205ms | 否 | ✓ 1697ms | 否 | ✓ 1783ms | http |
| 113.59.32.163:22222 | ✓ 1075ms | ✓ 1316ms | 否 | ✓ 1263ms | 否 | http |
| 138.124.90.140:1080 | ✓ 824ms | 否 | 否 | ✓ 1948ms | ✓ 1200ms | http |
| 95.3.9.78:3128 | ✓ 1203ms | ✓ 1892ms | ✓ 832ms | ✓ 1803ms | ✓ 1393ms | http |
| 91.107.141.42:8081 | ✓ 925ms | 否 | ✓ 1637ms | 否 | ✓ 1782ms | http |

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
