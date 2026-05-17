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

最后更新：2026-05-17 15:51:08 UTC（2026-05-17 23:51:08 UTC+8）

**代理总数：57**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1819ms | ✓ 1638ms | ✓ 1206ms | ✓ 1482ms | ✓ 1014ms | http |
| 115.231.181.40:8128 | ✓ 1325ms | ✓ 1850ms | ✓ 1632ms | 否 | ✓ 1031ms | http |
| 218.108.131.186:17890 | ✓ 878ms | ✓ 1140ms | ✓ 899ms | ✓ 1094ms | ✓ 922ms | http |
| 170.106.136.181:31002 | ✓ 1008ms | 否 | ✓ 624ms | ✓ 695ms | 否 | http |
| 185.200.188.234:10001 | ✓ 1127ms | 否 | ✓ 819ms | 否 | ✓ 1844ms | http |
| 59.46.216.131:30001 | ✓ 1765ms | ✓ 1391ms | ✓ 1705ms | 否 | ✓ 1131ms | http |
| 217.76.245.80:999 | ✓ 952ms | 否 | ✓ 1474ms | ✓ 1588ms | ✓ 1251ms | http |
| 91.242.229.129:8092 | ✓ 1925ms | 否 | ✓ 1944ms | 否 | ✓ 1950ms | http |
| 178.63.155.151:8898 | ✓ 1722ms | 否 | ✓ 1341ms | 否 | ✓ 1863ms | http |
| 129.80.217.21:444 | ✓ 374ms | 否 | ✓ 308ms | ✓ 1140ms | ✓ 879ms | http |
| 129.80.238.83:444 | ✓ 493ms | ✓ 1144ms | 否 | ✓ 1150ms | 否 | http |
| 8.154.21.175:3128 | ✓ 889ms | ✓ 1076ms | ✓ 934ms | ✓ 1155ms | ✓ 998ms | http |
| 128.199.121.61:9090 | ✓ 775ms | 否 | ✓ 1320ms | ✓ 1823ms | ✓ 1574ms | http |
| 128.199.113.85:9090 | ✓ 1952ms | 否 | ✓ 1743ms | ✓ 1235ms | ✓ 933ms | http |
| 148.230.4.241:999 | ✓ 955ms | 否 | ✓ 741ms | ✓ 1394ms | ✓ 1323ms | http |
| 129.154.225.163:8100 | ✓ 1407ms | 否 | 否 | ✓ 1928ms | ✓ 1284ms | http |
| 120.92.212.16:8890 | ✓ 905ms | ✓ 1329ms | ✓ 1331ms | ✓ 1875ms | ✓ 1554ms | http |
| 43.156.90.221:10808 | ✓ 1823ms | 否 | 否 | ✓ 1023ms | ✓ 1361ms | http |
| 5.252.33.13:2025 | ✓ 1620ms | 否 | ✓ 1997ms | 否 | ✓ 1877ms | http |
| 120.92.212.16:7890 | ✓ 1732ms | 否 | ✓ 1941ms | 否 | ✓ 1969ms | http |
| 128.199.114.189:9090 | ✓ 1492ms | 否 | ✓ 1141ms | ✓ 1519ms | 否 | http |
| 51.161.50.166:3128 | ✓ 1110ms | 否 | 否 | ✓ 1876ms | ✓ 1601ms | http |
| 103.21.220.141:3128 | ✓ 828ms | 否 | ✓ 761ms | ✓ 1017ms | ✓ 784ms | http |
| 47.112.25.109:7890 | ✓ 1977ms | 否 | ✓ 985ms | ✓ 1931ms | 否 | http |
| 185.21.15.206:3128 | ✓ 1513ms | 否 | ✓ 1865ms | ✓ 1971ms | 否 | http |
| 84.47.150.125:1080 | 否 | 否 | ✓ 1745ms | ✓ 1633ms | ✓ 1777ms | http |
| 210.223.44.230:3128 | ✓ 1343ms | 否 | ✓ 1340ms | ✓ 1710ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1786ms | 否 | ✓ 1293ms | ✓ 1503ms | ✓ 1093ms | http |
| 178.63.155.151:8888 | ✓ 1099ms | 否 | ✓ 1714ms | 否 | ✓ 1550ms | http |
| 133.18.123.225:26021 | 否 | 否 | ✓ 1037ms | ✓ 1180ms | ✓ 1125ms | http |
| 114.214.165.78:10810 | ✓ 1634ms | 否 | ✓ 1687ms | ✓ 1486ms | ✓ 1570ms | http |
| 152.70.91.193:40000 | ✓ 1649ms | 否 | 否 | ✓ 1721ms | ✓ 1584ms | http |
| 101.32.244.83:8080 | ✓ 954ms | 否 | ✓ 965ms | ✓ 1601ms | ✓ 1247ms | http |
| 121.43.196.213:8222 | ✓ 951ms | ✓ 1104ms | ✓ 863ms | ✓ 1115ms | ✓ 962ms | http |
| 121.43.196.210:8222 | ✓ 954ms | ✓ 1133ms | ✓ 823ms | ✓ 1136ms | ✓ 950ms | http |
| 166.88.55.83:7890 | ✓ 706ms | ✓ 1306ms | ✓ 769ms | ✓ 886ms | ✓ 702ms | http |
| 159.223.41.216:9090 | ✓ 758ms | 否 | ✓ 1551ms | ✓ 1070ms | ✓ 898ms | http |
| 86.104.72.220:1082 | ✓ 1322ms | 否 | ✓ 1069ms | ✓ 1876ms | 否 | http |
| 46.105.190.38:3128 | ✓ 1021ms | 否 | ✓ 1088ms | 否 | ✓ 1813ms | http |
| 3.101.133.120:80 | ✓ 506ms | 否 | ✓ 1416ms | ✓ 1215ms | ✓ 987ms | http |
| 114.214.170.41:27890 | ✓ 1191ms | ✓ 1440ms | ✓ 1330ms | ✓ 1468ms | ✓ 1184ms | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 1061ms | ✓ 977ms | ✓ 846ms | http |
| 20.27.14.220:8561 | ✓ 1806ms | ✓ 1218ms | ✓ 1110ms | ✓ 1282ms | ✓ 1176ms | http |
| 20.27.11.248:8561 | ✓ 1804ms | ✓ 1218ms | ✓ 1110ms | ✓ 1282ms | ✓ 1176ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1653ms | ✓ 1718ms | ✓ 1703ms | http |
| 45.153.231.229:8080 | ✓ 817ms | 否 | ✓ 1641ms | 否 | ✓ 1945ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1784ms | ✓ 1639ms | ✓ 1234ms | ✓ 1605ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1168ms | 否 | ✓ 845ms | ✓ 1088ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1056ms | ✓ 1212ms | ✓ 1407ms | http |
| 138.2.239.213:10010 | ✓ 758ms | 否 | 否 | ✓ 1401ms | ✓ 1398ms | http |
| 212.58.132.5:8888 | ✓ 1831ms | 否 | ✓ 1472ms | ✓ 1567ms | ✓ 1368ms | http |
| 47.74.226.8:5001 | ✓ 1847ms | ✓ 1388ms | ✓ 1089ms | ✓ 1312ms | ✓ 1245ms | http |
| 20.27.15.111:8561 | ✓ 1070ms | ✓ 1405ms | ✓ 1258ms | ✓ 1426ms | ✓ 1135ms | http |
| 61.52.131.172:8443 | 否 | ✓ 1806ms | ✓ 917ms | ✓ 1783ms | 否 | http |
| 103.172.70.173:8080 | ✓ 1822ms | 否 | 否 | ✓ 1534ms | ✓ 1614ms | http |
| 152.42.170.187:9090 | 否 | 否 | ✓ 1335ms | ✓ 1924ms | ✓ 916ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1930ms | ✓ 1848ms | ✓ 1857ms | http |

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
