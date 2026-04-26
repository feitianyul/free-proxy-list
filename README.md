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

最后更新：2026-04-26 10:48:22 UTC（2026-04-26 18:48:22 UTC+8）

**代理总数：37**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 37 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 37 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1270ms | ✓ 987ms | ✓ 969ms | ✓ 881ms | ✓ 674ms | http |
| 113.160.132.26:8080 | ✓ 1779ms | ✓ 1381ms | ✓ 1235ms | ✓ 1619ms | ✓ 1280ms | http |
| 34.96.238.40:8080 | ✓ 1624ms | ✓ 1661ms | ✓ 1807ms | 否 | 否 | http |
| 206.206.126.177:2412 | ✓ 1418ms | ✓ 1823ms | ✓ 1036ms | ✓ 947ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1872ms | 否 | 否 | ✓ 1837ms | ✓ 1703ms | http |
| 46.101.95.183:8888 | ✓ 957ms | 否 | ✓ 1181ms | ✓ 1712ms | ✓ 1397ms | http |
| 103.187.146.151:3128 | ✓ 1247ms | 否 | ✓ 951ms | ✓ 1375ms | ✓ 947ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1303ms | ✓ 1028ms | ✓ 1309ms | ✓ 1060ms | http |
| 47.84.73.61:1080 | ✓ 1417ms | ✓ 1658ms | ✓ 849ms | 否 | 否 | http |
| 16.163.173.29:1080 | ✓ 615ms | ✓ 1213ms | ✓ 1108ms | ✓ 792ms | ✓ 624ms | http |
| 36.141.21.200:7890 | 否 | ✓ 1126ms | ✓ 861ms | ✓ 1175ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1515ms | 否 | 否 | ✓ 1137ms | ✓ 924ms | http |
| 130.61.139.145:3128 | ✓ 1286ms | 否 | ✓ 1953ms | ✓ 1811ms | ✓ 1479ms | http |
| 2.27.54.161:1080 | ✓ 1326ms | 否 | ✓ 1418ms | 否 | ✓ 1593ms | http |
| 47.85.51.197:1080 | ✓ 283ms | ✓ 1211ms | ✓ 1805ms | 否 | ✓ 1224ms | http |
| 120.92.212.16:7890 | ✓ 1691ms | 否 | 否 | ✓ 1395ms | ✓ 1929ms | http |
| 45.76.207.177:40000 | 否 | 否 | ✓ 1998ms | ✓ 1426ms | ✓ 903ms | http |
| 210.223.44.230:3128 | ✓ 1349ms | ✓ 1323ms | ✓ 1320ms | 否 | ✓ 1325ms | http |
| 86.109.3.28:10041 | ✓ 1037ms | 否 | 否 | ✓ 1826ms | ✓ 1114ms | http |
| 8.219.195.129:1080 | ✓ 1456ms | ✓ 1632ms | ✓ 762ms | ✓ 1028ms | ✓ 807ms | http |
| 152.42.177.32:8888 | ✓ 898ms | 否 | ✓ 934ms | ✓ 1227ms | 否 | http |
| 177.93.132.244:3128 | ✓ 953ms | 否 | ✓ 889ms | 否 | ✓ 1828ms | http |
| 159.89.31.62:8080 | ✓ 1157ms | 否 | ✓ 1952ms | 否 | ✓ 1669ms | http |
| 103.167.169.22:8080 | ✓ 1034ms | 否 | ✓ 1158ms | ✓ 1358ms | 否 | http |
| 168.144.75.9:3128 | 否 | 否 | ✓ 1927ms | ✓ 1988ms | ✓ 1343ms | http |
| 43.133.90.161:8888 | 否 | ✓ 1730ms | 否 | ✓ 1507ms | ✓ 1908ms | http |
| 218.108.131.186:17890 | ✓ 774ms | ✓ 994ms | ✓ 818ms | ✓ 1077ms | ✓ 853ms | http |
| 91.99.15.45:2095 | ✓ 1297ms | ✓ 1768ms | 否 | ✓ 1991ms | 否 | http |
| 20.120.225.109:3128 | ✓ 952ms | ✓ 1467ms | ✓ 1207ms | ✓ 1096ms | ✓ 824ms | http |
| 62.113.119.14:8080 | ✓ 924ms | 否 | ✓ 1566ms | ✓ 1841ms | ✓ 1261ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1663ms | ✓ 1584ms | ✓ 1295ms | http |
| 47.101.159.19:8899 | ✓ 811ms | ✓ 981ms | ✓ 856ms | ✓ 1080ms | ✓ 888ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1149ms | ✓ 1243ms | ✓ 952ms | http |
| 124.16.93.233:7890 | ✓ 898ms | ✓ 1079ms | ✓ 1929ms | ✓ 1456ms | ✓ 955ms | http |
| 61.52.131.172:8443 | ✓ 862ms | ✓ 1168ms | ✓ 974ms | ✓ 1257ms | ✓ 998ms | http |
| 117.84.29.218:1080 | ✓ 1638ms | 否 | ✓ 1269ms | ✓ 1880ms | 否 | http |
| 103.39.51.207:8080 | ✓ 1260ms | 否 | ✓ 1668ms | ✓ 1536ms | ✓ 1679ms | http |

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
