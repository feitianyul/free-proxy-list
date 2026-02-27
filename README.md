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

最后更新：2026-02-27 15:37:20 UTC（2026-02-27 23:37:20 UTC+8）

**代理总数：61**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 353ms | ✓ 1085ms | ✓ 1686ms | ✓ 1208ms | ✓ 1123ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1882ms | ✓ 1934ms | ✓ 1117ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1272ms | ✓ 1126ms | ✓ 1051ms | http |
| 52.188.28.218:3128 | ✓ 204ms | 否 | ✓ 186ms | 否 | ✓ 727ms | http |
| 34.142.0.1:10808 | ✓ 532ms | 否 | ✓ 1875ms | ✓ 1697ms | ✓ 1699ms | http |
| 132.145.93.138:1080 | ✓ 853ms | 否 | ✓ 1027ms | ✓ 1726ms | 否 | http |
| 14.143.222.113:10158 | ✓ 1098ms | 否 | ✓ 1098ms | ✓ 1464ms | 否 | http |
| 45.88.0.111:3128 | ✓ 887ms | 否 | ✓ 1252ms | ✓ 1733ms | 否 | http |
| 35.234.17.221:8080 | ✓ 1817ms | ✓ 1460ms | 否 | ✓ 1180ms | 否 | http |
| 85.208.108.43:2094 | ✓ 189ms | 否 | ✓ 711ms | ✓ 1161ms | ✓ 764ms | http |
| 72.56.59.62:63133 | ✓ 1649ms | 否 | ✓ 1807ms | 否 | ✓ 1997ms | http |
| 72.56.59.56:63127 | ✓ 1639ms | 否 | ✓ 1780ms | 否 | ✓ 1997ms | http |
| 81.70.169.194:80 | ✓ 1050ms | ✓ 1756ms | 否 | 否 | ✓ 1101ms | http |
| 195.123.209.48:3128 | ✓ 1155ms | 否 | ✓ 954ms | ✓ 1781ms | ✓ 1412ms | http |
| 120.92.212.16:7890 | ✓ 1178ms | 否 | ✓ 1277ms | ✓ 1312ms | 否 | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1539ms | ✓ 1746ms | ✓ 1092ms | http |
| 201.150.116.32:999 | ✓ 700ms | 否 | ✓ 932ms | 否 | ✓ 1036ms | http |
| 120.46.152.136:3128 | ✓ 979ms | ✓ 1223ms | ✓ 1047ms | ✓ 1261ms | ✓ 980ms | http |
| 94.177.131.12:3128 | ✓ 536ms | 否 | ✓ 537ms | ✓ 906ms | ✓ 702ms | http |
| 91.238.105.64:2024 | ✓ 976ms | 否 | ✓ 771ms | ✓ 1634ms | ✓ 1171ms | http |
| 103.117.100.127:13082 | ✓ 751ms | 否 | ✓ 966ms | ✓ 930ms | ✓ 772ms | http |
| 91.238.104.172:2024 | ✓ 674ms | 否 | ✓ 1062ms | ✓ 1628ms | 否 | http |
| 61.72.110.24:3128 | ✓ 1782ms | ✓ 1201ms | 否 | 否 | ✓ 1477ms | http |
| 91.238.104.171:2023 | ✓ 987ms | 否 | ✓ 894ms | ✓ 1652ms | 否 | http |
| 20.210.39.153:8561 | ✓ 881ms | 否 | ✓ 1705ms | ✓ 1999ms | ✓ 1763ms | http |
| 20.78.118.91:8561 | ✓ 704ms | 否 | ✓ 1647ms | ✓ 1978ms | ✓ 1764ms | http |
| 20.78.26.206:8561 | ✓ 1271ms | ✓ 1877ms | ✓ 1702ms | ✓ 1983ms | ✓ 1989ms | http |
| 170.78.208.245:999 | 否 | 否 | ✓ 1260ms | ✓ 1083ms | ✓ 947ms | http |
| 45.140.147.155:1081 | ✓ 597ms | ✓ 1761ms | ✓ 709ms | ✓ 1779ms | 否 | http |
| 42.115.72.179:2033 | ✓ 1729ms | 否 | ✓ 1561ms | ✓ 1861ms | 否 | http |
| 138.124.53.25:7443 | ✓ 977ms | 否 | 否 | ✓ 1799ms | ✓ 1725ms | http |
| 72.56.50.17:59787 | ✓ 1657ms | 否 | ✓ 1753ms | 否 | ✓ 1999ms | http |
| 103.84.95.54:7890 | ✓ 1612ms | 否 | ✓ 823ms | ✓ 1940ms | 否 | http |
| 61.72.110.94:3128 | ✓ 1690ms | 否 | ✓ 1240ms | ✓ 1589ms | ✓ 1866ms | http |
| 147.45.216.148:1080 | ✓ 520ms | 否 | ✓ 1780ms | 否 | ✓ 1289ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1132ms | 否 | ✓ 1176ms | ✓ 922ms | http |
| 120.92.212.16:8890 | ✓ 1022ms | ✓ 1393ms | 否 | ✓ 1352ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1506ms | 否 | ✓ 1419ms | ✓ 1972ms | ✓ 1342ms | http |
| 14.56.107.244:3128 | ✓ 1598ms | 否 | ✓ 1965ms | ✓ 1096ms | 否 | http |
| 199.68.217.2:3128 | ✓ 345ms | ✓ 951ms | ✓ 1128ms | ✓ 1016ms | ✓ 934ms | http |
| 121.237.181.137:8888 | ✓ 944ms | ✓ 1295ms | ✓ 917ms | 否 | ✓ 1951ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1430ms | ✓ 1309ms | 否 | ✓ 1992ms | http |
| 168.235.110.63:3128 | 否 | 否 | ✓ 1289ms | ✓ 1190ms | ✓ 991ms | http |
| 39.104.201.40:7890 | ✓ 988ms | ✓ 1299ms | ✓ 1067ms | ✓ 1281ms | ✓ 1052ms | http |
| 120.132.97.88:7897 | ✓ 926ms | 否 | ✓ 1116ms | ✓ 1254ms | 否 | http |
| 64.181.240.152:3128 | ✓ 846ms | 否 | ✓ 1668ms | 否 | ✓ 1108ms | http |
| 45.136.198.40:3128 | 否 | 否 | ✓ 691ms | ✓ 1533ms | ✓ 1250ms | http |
| 150.107.140.238:3128 | ✓ 1927ms | 否 | ✓ 1686ms | 否 | ✓ 1470ms | http |
| 103.215.36.88:16316 | ✓ 1073ms | 否 | ✓ 1127ms | ✓ 1580ms | 否 | http |
| 36.147.78.166:80 | 否 | 否 | ✓ 1771ms | ✓ 1983ms | ✓ 1684ms | http |
| 121.147.253.205:3096 | ✓ 1775ms | ✓ 1598ms | ✓ 1986ms | ✓ 1177ms | ✓ 911ms | http |
| 212.227.189.168:3128 | ✓ 755ms | 否 | ✓ 1364ms | 否 | ✓ 1473ms | http |
| 101.43.255.96:80 | ✓ 1018ms | ✓ 1408ms | ✓ 1467ms | 否 | 否 | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1319ms | ✓ 1650ms | ✓ 1605ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1469ms | ✓ 1947ms | ✓ 1853ms | 否 | http |
| 116.80.64.44:7777 | ✓ 1751ms | 否 | ✓ 1929ms | 否 | ✓ 1743ms | http |
| 104.37.184.214:1080 | 否 | 否 | ✓ 1758ms | ✓ 1031ms | ✓ 1386ms | http |
| 172.212.68.37:3128 | ✓ 292ms | ✓ 1856ms | ✓ 1104ms | ✓ 1604ms | ✓ 975ms | http |
| 120.55.163.237:10086 | ✓ 1344ms | ✓ 1354ms | ✓ 1447ms | 否 | 否 | http |
| 125.128.12.34:3128 | ✓ 814ms | ✓ 1275ms | ✓ 1378ms | 否 | 否 | http |
| 220.197.44.36:3128 | ✓ 1852ms | 否 | ✓ 1519ms | 否 | ✓ 1729ms | http |

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
